package e4test

import (
	"crypto/rand"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	e4 "teserakt/e4go/pkg/e4common"
	"time"
)

// TestIDKey provides a representation of an id-key pair.
type TestIDKey struct {
	ID  []byte
	Key []byte
}

// TestTopicKey provides a representation of the topic-key pair.
type TestTopicKey struct {
	TopicName string
	Key       []byte
}

// TestResult reports a test outcome for nice display
type TestResult struct {
	Name     string
	Result   bool
	Critical bool
	Error    error
}

// New generates a new TestIDKey
func (t *TestIDKey) New() error {
	id, err := GenerateID()
	if err != nil {
		return err
	}
	key, err := GenerateKey()
	if err != nil {
		return err
	}
	t.ID = id
	t.Key = key
	return nil
}

// GetHexID returns a hex-encoded Identifier
func (t *TestIDKey) GetHexID() string {
	return hex.EncodeToString(t.ID)
}

// GetHexKey returns a hex-encoded Key
func (t *TestIDKey) GetHexKey() string {
	return hex.EncodeToString(t.Key)
}

// New generates a new TestIDKey
func (t *TestTopicKey) New() error {
	topic, err := GenerateTopic()
	if err != nil {
		return err
	}
	key, err := GenerateKey()
	if err != nil {
		return err
	}
	t.TopicName = topic
	t.Key = key
	return nil
}

// GetHexKey returns a hex-encoded Key
func (t *TestTopicKey) GetHexKey() string {
	return hex.EncodeToString(t.Key)
}

// FindAndCheckPathFile does some light sanity checks on
// file access. If supplied an absolute file path, it checks
// we can stat this file and that it isn't stupid, like
// a directory
// if the subpath variable is a relative path this is assumed
// to be relative to one of the directories specified
// as a gopath. We search each one to try to find the
// file and return its absolute path.
func FindAndCheckPathFile(subpath string) (string, error) {

	// if we
	if filepath.IsAbs(subpath) {
		fileInfo, err := os.Stat(subpath)
		if err != nil {
			return "", fmt.Errorf("Unable to stat file %s", subpath)
		}
		if fileInfo.IsDir() {
			return "", fmt.Errorf("Can't exec a directory %s", subpath)
		}
		return subpath, nil
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	fullfilepath := ""

	godirs := strings.Split(gopath, ":")
	for _, godir := range godirs {
		fileTentative := filepath.Join(godir, subpath)
		fileInfo, err := os.Stat(fileTentative)
		if err == nil && !fileInfo.IsDir() {
			fullfilepath = string(fileTentative)
			break
		}
	}
	if fullfilepath == "" {
		return "", fmt.Errorf("Unable to locate file %s in any Gopath directory", subpath)
	}
	return fullfilepath, nil
}

// RunDaemon launches the specified process with arguments
// and waits for notification on an error channel to
// signal the process to exit. Should be used for testing
// daemons like the C2 backend, where the process needs to
// stay alive and we interact with it via an API.
// **Note**: procenv is currently appended to os.Environ(),
// no merging takes place at all.
func RunDaemon(errc chan *TestResult,
	stopc chan struct{},
	clientwaitc chan struct{},
	path string,
	args []string,
	procenv []string) {

	osenv := os.Environ()

	// Not quite correct
	procenv = append(procenv, osenv...)

	subproc := exec.Cmd{
		Path: path,
		Args: args,
		Env:  procenv,
	}

	/*spstdin, _ := subproc.StdinPipe()*/
	spstdout, _ := subproc.StdoutPipe()
	spstderr, _ := subproc.StderrPipe()

	if err := subproc.Start(); err != nil {
		wrappederr := fmt.Sprintf("runDaemon failed. %s", err)
		fmt.Fprintf(os.Stdout, wrappederr)
		close(clientwaitc)
		errc <- &TestResult{
			Name:     "",
			Result:   false,
			Critical: true,
			Error:    errors.New(wrappederr),
		}
		return
	}

	fmt.Fprintf(os.Stdout, "Running %s as %d\n", path, subproc.Process.Pid)

	// clean up on exit:
	defer func() {
		// send an interrupt signal to terminate the process.
		subproc.Wait()
		close(clientwaitc)
	}()

	// tell the caller we've set up correctly:
	time.Sleep(200 * time.Millisecond)
	clientwaitc <- struct{}{}
	// wait for signal on stop channel:

	<-stopc
	subproc.Process.Signal(os.Interrupt)

	bytes, _ := ioutil.ReadAll(spstdout)
	os.Stdout.Write(bytes)
	bytes, _ = ioutil.ReadAll(spstderr)
	os.Stdout.Write(bytes)
}

// RunCommand launches the specified process with arguments
// and waits for it to exit, returning the contents of stdout and stderr
func RunCommand(path string,
	args []string,
	procenv []string) ([]byte, []byte, error) {

	env := os.Environ()

	// Not quite correct
	procenv = append(procenv, env...)

	subproc := exec.Cmd{
		Path: path,
		Args: args,
		Env:  env,
	}

	/*spstdin, _ := subproc.StdinPipe()*/
	spstdout, _ := subproc.StdoutPipe()
	spstderr, _ := subproc.StderrPipe()

	if err := subproc.Run(); err != nil {
		return nil, nil, fmt.Errorf("runDaemon failed. %s", err)
	}

	subproc.Wait()
	stdoutbytes, err := ioutil.ReadAll(spstdout)
	if err != nil {
		return nil, nil, err
	}
	stderrbytes, err := ioutil.ReadAll(spstderr)
	if err != nil {
		return nil, nil, err
	}
	return stdoutbytes, stderrbytes, nil
}

// GetRandomDBName produces a random database
// path in tmp for use with SQLite3 testing.
func GetRandomDBName() string {
	bytes := [16]byte{}
	_, err := rand.Read(bytes[:])
	if err != nil {
		panic(err)
	}
	dbCandidate := b64.StdEncoding.EncodeToString(bytes[:])
	dbCleaned1 := strings.Replace(dbCandidate, "+", "", -1)
	dbCleaned2 := strings.Replace(dbCleaned1, "/", "", -1)
	dbCleaned3 := strings.Replace(dbCleaned2, "=", "", -1)

	dbPath := fmt.Sprintf("/tmp/e4c2_unittest_%s.sqlite", dbCleaned3)
	return dbPath
}

// GenerateID generates a random ID that is e4.IDLen bytes
// in length, using a CSPRNG
func GenerateID() ([]byte, error) {
	idbytes := [e4.IDLen]byte{}
	_, err := rand.Read(idbytes[:])
	if err != nil {
		return nil, err
	}
	return idbytes[:], nil
}

// GenerateKey generates a random key that is e4.KeyLen bytes
// in length, using a CSPRNG
func GenerateKey() ([]byte, error) {
	keybytes := [e4.KeyLen]byte{}
	_, err := rand.Read(keybytes[:])
	if err != nil {
		return nil, err
	}
	return keybytes[:], nil
}

// GenerateTopic generates a random topic
func GenerateTopic() (string, error) {
	bytes := [28]byte{}
	_, err := rand.Read(bytes[:])
	if err != nil {
		return "", err
	}
	tCandidate := b64.StdEncoding.EncodeToString(bytes[:])
	tCleaned1 := strings.Replace(tCandidate, "+", "", -1)
	tCleaned2 := strings.Replace(tCleaned1, "/", "", -1)
	tCleaned3 := strings.Replace(tCleaned2, "=", "", -1)
	return tCleaned3[0:32], nil
}

// ConstructHTTPSClient builds an HTTPS client for use
// with the API
func ConstructHTTPSClient() http.Client {
	tlsConfig := &tls.Config{
		//Certificates: []tls.Certificate{tlsCert},
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		InsecureSkipVerify: true,
	}

	httpTransport := &http.Transport{
		IdleConnTimeout: time.Second * 20, // required, or goroutines can hang.
		TLSClientConfig: tlsConfig,
		//TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	httpClient := http.Client{
		Transport: httpTransport,
	}

	return httpClient
}