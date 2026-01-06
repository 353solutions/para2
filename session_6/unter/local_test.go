package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
- build executable
- run it
- run tests over it
- kill the executable

Issues:
- Don't use "go run", the process ID is for the "go run", not the server
- Tests can run in parallel - which port to use?
*/

func freePort(t *testing.T) int {
	lis, err := net.Listen("tcp", "")
	require.NoError(t, err, "listen")
	lis.Close()

	return lis.Addr().(*net.TCPAddr).Port
}

func buildServer(t *testing.T) string {
	exe := path.Join(t.TempDir(), "unter")
	cmd := exec.Command("go", "build", "-o", exe, ".")
	err := cmd.Run()
	if err != nil {
		t.Fatal(err)
	}

	return exe
}

func runServer(t *testing.T) int {
	exe := buildServer(t)
	port := freePort(t)

	cmd := exec.Command(exe)
	cmd.Env = append(os.Environ(), fmt.Sprintf("UNTER_ADDR=:%d", port))
	cmd.Env = append(os.Environ(), fmt.Sprintf("UNTER_DB=%s", path.Join(t.TempDir(), "unter.db")))

	err := cmd.Start()
	require.NoError(t, err, "start")
	t.Cleanup(func() {
		if err := cmd.Process.Kill(); err != nil {
			t.Logf("kill %d: %v", cmd.Process.Pid, err)
		}
	})

	return port
}

func TestHealth_Local(t *testing.T) {
	port := runServer(t)
	url := fmt.Sprintf("http://localhost:%d/health", port)
	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, url, nil)
	require.NoError(t, err, "request")

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err, "get")
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
