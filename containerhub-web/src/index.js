import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';
import './styles/style.css';
import '@xterm/xterm/css/xterm.css';
import 'bootstrap-icons/font/bootstrap-icons.css';
import ContainerPage from './pages/container';

import m from 'mithril';
import LoginPage from './pages/login';
import { currentUser } from './api';
import ImagePage from './pages/image';
import KeysPage from './pages/keys';
import SSHPage from './pages/ssh';
import DocumentationPage from './pages/documentation';

function guard(component) {
  return {
    onmatch() {
      if (currentUser()) {
        return component;
      }
      m.route.set('/login');
    },
  };
}

m.route(document.body, '/', {
  '/': {
    onmatch: () => {
      m.route.set('/containers');
    },
  },
  '/containers': guard(ContainerPage),
  '/images': ImagePage,
  '/login': LoginPage,
  '/docs': DocumentationPage,
  '/keys': guard(KeysPage),
  '/ssh': guard(SSHPage),
});
