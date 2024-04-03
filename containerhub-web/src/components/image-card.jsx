import { createContainer, currentUser } from '../api';

function ImageCard({ attrs: { image } }) {
  let customName = '';

  function handleCreate() {
    if (!currentUser()) {
      m.route.set('/login');
      return;
    }
    if (!customName) {
      return;
    }
    createContainer(image.name, customName).then(() => {
      m.route.set('/containers');
    });
  }

  return {
    view() {
      return (
        <div className="card shadow">
          <h5 class="card-header">{image.name}</h5>
          <div class="card-body">
            <dl class="row mb-0">
              <dt class="col-sm-3">Name</dt>
              <dd class="col-sm-9">{image.imageName}</dd>
              <dt class="col-sm-3">Author</dt>
              <dd class="col-sm-9">{image.author}</dd>
              <dt class="col-sm-3">Tags</dt>
              <dd class="col-sm-9">
                {image.tags.map((tag) => (
                  <span class="badge rounded-pill text-bg-secondary me-1">{tag}</span>
                ))}
              </dd>
            </dl>
            <p class="cart-text mt-2">{image.description}</p>
            <div class="gap-2 d-flex mt-3 ">
              <button class="btn btn-sm btn-primary" data-bs-toggle="modal" data-bs-target="#createModal">
                Create
              </button>
            </div>
          </div>
          <div class="modal fade" id="createModal">
            <div class="modal-dialog modal-dialog-centered">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">Create New Container</h5>
                </div>
                <div class="modal-body">
                  <div>
                    <label htmlFor="custom-name" className="form-label">
                      Custom Name
                    </label>
                    <input
                      id="custom-name"
                      class="form-control"
                      type="text"
                      value={customName}
                      oninput={(e) => (customName = e.target.value)}
                    />
                  </div>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">
                    Close
                  </button>
                  <button
                    disabled={!customName}
                    type="button"
                    class="btn btn-primary"
                    onclick={() => handleCreate()}
                    data-bs-dismiss="modal"
                  >
                    Create
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

export default ImageCard;
