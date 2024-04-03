import classNames from 'classnames';
import { startContainer, stopContainer, destroyContainer } from '../api';

function ContainerCard({ attrs: { container, onchange } }) {
  let actionRunning = false;

  const actions = {
    start: startContainer,
    stop: stopContainer,
    destroy: destroyContainer,
  };

  function handleAction(action) {
    actionRunning = true;
    actions[action](container.Id)
      .then(() => {
        onchange();
        actionRunning = false;
      })
      .catch((error) => {
        console.error(error);
      });
  }

  return {
    view({ attrs: { container } }) {
      const d = new Date(container.Created * 1000);
      return (
        <div class="card shadow">
          <h5 class="card-header">{container.Labels["containerhub-name"]}</h5>
          <div class="card-body">
            <dl class="row mb-0">
              <dt class="col-sm-3">ID</dt>
              <dd class="col-sm-9 font-monospace fw-bold">{container.Id.substring(0, 12)}</dd>
              <dt class="col-sm-3">Image</dt>
              <dd class="col-sm-9">{container.Image}</dd>
              <dt class="col-sm-3">State</dt>
              <dd class="col-sm-9">
                <span
                  class={classNames('badge', {
                    'text-bg-success': container.State === 'running',
                    'text-bg-secondary': container.State !== 'running',
                  })}
                >
                  {container.State}
                </span>
              </dd>
              <dt class="col-sm-3">Status</dt>
              <dd class="col-sm-9">{container.Status}</dd>
              <dt class="col-sm-3">Created</dt>
              <dd class="col-sm-9">
                {d.toLocaleDateString()} {d.toLocaleTimeString()}
              </dd>
            </dl>
            <div class="gap-2 d-flex mt-3">
              <button
                class="btn btn-primary"
                onclick={() => handleAction(container.State === 'running' ? 'stop' : 'start')}
                disabled={actionRunning}
              >
                {actionRunning && <span class="spinner-border spinner-border-sm me-1"></span>}
                <span role="status">{container.State === 'running' ? 'Stop' : 'Start'}</span>
              </button>
              <button class="btn btn-danger" data-bs-toggle="modal" data-bs-target="#destroyModal">
                Destroy
              </button>
            </div>
          </div>
          <div class="modal fade" id="destroyModal">
            <div class="modal-dialog modal-dialog-centered">
              <div class="modal-content">
                <div class="modal-body">
                  <p>Are you sure you want to destroy this container?</p>
                  <div className="alert alert-danger">
                    <strong>Warning!</strong> This action cannot be undone.
                  </div>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">
                    Close
                  </button>
                  <button
                    type="button"
                    class="btn btn-danger"
                    onclick={() => handleAction('destroy')}
                    data-bs-dismiss="modal"
                  >
                    Destroy
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      );
    },
  };
}

export default ContainerCard;
