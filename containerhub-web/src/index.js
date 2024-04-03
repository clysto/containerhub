import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';
import './styles/style.css';
import ContainerPage from './pages/container';

import m from 'mithril';
import LoginPage from './pages/login';
import { currentUser } from './api';
import ImagePage from './pages/image';
import KeysPage from './pages/keys';

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
  '/keys': guard(KeysPage),
});
