import './styles/style.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap';
import ContainerPage from './pages/container';

import m from 'mithril';
import LoginPage from './pages/login';

m.route(document.body, '/', {
  '/': {
    onmatch: () => {
      m.route.set('/containers');
    },
  },
  '/containers': ContainerPage,
  '/login': LoginPage,
});
