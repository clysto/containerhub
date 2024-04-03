import Layout from '../components/layout';
import { listContainers } from '../api';

function ContainerPage() {
  let containers = [];
  return {
    oninit() {
      listContainers().then((data) => {
        containers = data;
      });
    },
    view() {
      return (
        <Layout>
          <div className="container p-4">
            {containers.map((container) => (
              <div class="card">
                <div class="card-header">
                  <ul class="nav nav-tabs card-header-tabs">
                    <li class="nav-item">
                      <span class="nav-link active">Information</span>
                    </li>
                    <li class="nav-item">
                      <span class="nav-link">SSH Key</span>
                    </li>
                  </ul>
                </div>
                <div class="card-body">
                  <ul class="list-group list-group-flush font-monospace">
                    <li class="list-group-item">
                      <dl class="row mb-0">
                        <dt class="col-sm-3">Name</dt>
                        <dd class="col-sm-9">{container.Names[0]}</dd>
                        <dt class="col-sm-3">ID</dt>
                        <dd class="col-sm-9">{container.Id}</dd>
                        <dt class="col-sm-3">Image</dt>
                        <dd class="col-sm-9">{container.Image}</dd>
                        <dt class="col-sm-3">State</dt>
                        <dd class="col-sm-9">{container.State}</dd>
                        <dt class="col-sm-3">Status</dt>
                        <dd class="col-sm-9">{container.Status}</dd>
                      </dl>
                    </li>
                  </ul>

                  <div class="gap-2 d-flex mt-2">
                    <button class="btn btn-primary">Start</button>
                    <button class="btn btn-primary">Stop</button>
                    <button class="btn btn-danger">Destroy</button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </Layout>
      );
    },
  };
}

export default ContainerPage;
