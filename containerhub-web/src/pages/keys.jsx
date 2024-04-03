import { listContainers, listKeys } from '../api';
import Layout from '../components/layout';
import { generateKey, deleteKey } from '../api';

function textToOctetStreamURL(text) {
  const blob = new Blob([text], { type: 'application/octet-stream' });
  return URL.createObjectURL(blob);
}

function KeysPage() {
  let containers = [];
  let sshKeys = [];
  function handleGenerateKey() {
    const containerId = document.getElementById('container-id').value;
    if (containerId === '0') {
      return;
    }
    generateKey(containerId).then((data) => {
      sshKeys.push(data);
    });
  }
  function handleDeleteKey(containerID) {
    deleteKey(containerID).then(() => {
      sshKeys = sshKeys.filter((key) => key.containerID !== containerID);
    });
  }
  return {
    oninit() {
      listContainers().then((data) => {
        containers = data;
      });
      listKeys().then((data) => {
        sshKeys = data;
      });
    },
    view() {
      return (
        <Layout>
          <div class="border-bottom">
            <div className="container py-5">
              <h1>SSH Keys</h1>
              <p className="lead mb-0">Generate SSH keys for your containers.</p>
            </div>
          </div>
          <div className="container py-4">
            <div className="input-group">
              <select class="form-select" id="container-id">
                <option value="0" selected>
                  Choose a container...
                </option>
                {containers.map((container) => (
                  <option value={container.Id}>{container.Labels['containerhub-name']}</option>
                ))}
              </select>
              <button class="col-auto btn btn-primary" onclick={handleGenerateKey}>
                <i class="bi bi-key-fill"></i> Generate Key
              </button>
            </div>

            <table class="table mt-4">
              <thead>
                <tr>
                  <th scope="col">Container</th>
                  <th scope="col">Public Key</th>
                  <th scope="col"></th>
                </tr>
              </thead>
              <tbody>
                {sshKeys.map((key) => (
                  <tr>
                    <td>{key.containerName}</td>
                    <td>
                      <div class="font-monospace text-break">{key.publicKey}</div>
                    </td>
                    <td style="text-align: right">
                      <div class="btn-group">
                        <a
                          download={key.containerName + '.pem'}
                          className="btn btn-primary btn-sm"
                          href={textToOctetStreamURL(key.privateKey)}
                        >
                          <i class="bi bi-file-earmark-arrow-down-fill"></i>
                        </a>
                        <button className="btn btn-danger btn-sm" onclick={() => handleDeleteKey(key.containerID)}>
                          <i class="bi bi-trash-fill"></i>
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </Layout>
      );
    },
  };
}

export default KeysPage;
