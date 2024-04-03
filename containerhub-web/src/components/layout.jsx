import { Link } from 'mithril/route';

function Layout() {
  return {
    view(vnode) {
      return (
        <div>
          <nav class="navbar navbar-expand-lg bg-body-tertiary shadow-sm">
            <div class="container-fluid">
              <span class="navbar-brand">Container Hub</span>
              <button
                class="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarSupportedContent"
              >
                <span class="navbar-toggler-icon"></span>
              </button>
              <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                  <li class="nav-item">
                    <Link class="nav-link" href="/containers">
                      Containers
                    </Link>
                  </li>
                  <li class="nav-item">
                    <Link class="nav-link" href="#">
                      Others
                    </Link>
                  </li>
                </ul>
                <div class="d-flex">
                  <Link class="btn btn-primary block" href="/login">
                    Login
                  </Link>
                </div>
              </div>
            </div>
          </nav>
          <main>{vnode.children}</main>
        </div>
      );
    },
  };
}

export default Layout;
