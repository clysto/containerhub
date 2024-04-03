import { Link } from 'mithril/route';
import { currentUser, logout } from '../api';
import classNames from 'classnames';

function Layout() {
  const user = currentUser();
  function handleLogout() {
    logout();
    m.route.set('/login');
  }
  return {
    view(vnode) {
      return (
        <div class="d-flex flex-column min-vh-100">
          <nav class="navbar navbar-expand-lg shadow-sm bg-primary sticky-top" data-bs-theme="dark">
            <div class="container">
              <strong class="navbar-brand">Container Hub</strong>
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
                    <Link
                      class={classNames('nav-link', {
                        active: m.route.get().startsWith('/containers'),
                      })}
                      href="/containers"
                    >
                      Containers
                    </Link>
                  </li>
                  <li class="nav-item">
                    <Link
                      class={classNames('nav-link', {
                        active: m.route.get().startsWith('/keys'),
                      })}
                      href="/keys"
                    >
                      SSH Keys
                    </Link>
                  </li>
                  <li class="nav-item">
                    <Link
                      class={classNames('nav-link', {
                        active: m.route.get().startsWith('/images'),
                      })}
                      href="/images"
                    >
                      Images
                    </Link>
                  </li>
                  <li class="nav-item">
                    <Link
                      class={classNames('nav-link', {
                        active: m.route.get().startsWith('/docs'),
                      })}
                      href="/docs"
                    >
                      Documentation
                    </Link>
                  </li>
                </ul>
                <div class="d-flex">
                  {user ? (
                    <div class="dropdown nav-item" data-bs-theme="light">
                      <button class="btn btn-primary dropdown-toggle" data-bs-toggle="dropdown" aria-expanded="false">
                        {user.username}
                      </button>
                      <ul class="dropdown-menu dropdown-menu-end">
                        <li>
                          <button class="dropdown-item" onclick={handleLogout}>
                            <i class="bi bi-door-closed-fill"></i> Logout
                          </button>
                        </li>
                      </ul>
                    </div>
                  ) : (
                    <Link class="btn btn-outline-light" href="/login">
                      <i class="bi bi-door-open-fill"></i> Login
                    </Link>
                  )}
                </div>
              </div>
            </div>
          </nav>
          <main class="flex-grow-1">{vnode.children}</main>
          <footer class="footer py-4 py-md-5 bg-body-tertiary">
            <div className="container text-muted">Container Hub Interface, Copyright © 2023-2024</div>
          </footer>
        </div>
      );
    },
  };
}

export default Layout;
