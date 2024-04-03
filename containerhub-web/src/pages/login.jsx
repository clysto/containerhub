import Layout from '../components/layout';
import styles from '../styles/login.module.css';
import classNames from 'classnames';
import { login } from '../api';

function LoginPage() {
  let message = '';

  async function handleLogin() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    try {
      await login(username, password);
      m.route.set('/containers');
    } catch (error) {
      message = error.response.error;
    }
  }

  return {
    view() {
      return (
        <Layout>
          <div class="container py-4">
            <div class={classNames(styles.form, 'rounded-3 border shadow')}>
              <h1 class={styles.title}>Login</h1>
              <div class="mb-2">
                <label class="form-label" htmlFor="username">
                  Username
                </label>
                <input id="username" class="form-control" type="text" />
              </div>
              <div>
                <label class="form-label" htmlFor="password">
                  Password
                </label>
                <input id="password" class="form-control" type="password" />
              </div>
              {message && <div class="alert alert-warning mt-4">{message}</div>}
              <div class="mt-5 d-grid gap-2">
                <button class="btn btn-primary" onclick={handleLogin}>
                  <i class="bi bi-door-open-fill"></i> Login
                </button>
              </div>
            </div>
          </div>
        </Layout>
      );
    },
  };
}

export default LoginPage;
