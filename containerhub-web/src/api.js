import m from 'mithril';

const store = {
  user: null,
  token: null,
};

init();

function init() {
  const data = localStorage.getItem('store');
  if (data) {
    Object.assign(store, JSON.parse(data));
  }
}

function saveStore() {
  localStorage.setItem('store', JSON.stringify(store));
}

export async function login(username, password) {
  try {
    const response = await m.request('/api/v1/login', {
      method: 'POST',
      url: '/api/v1/login',
      body: { username, password },
    });
    store.token = response.token;
    store.user = response.user;
    saveStore();
    return response;
  } catch (error) {
    throw error;
  }
}

export function logout() {
  store.token = null;
  store.user = null;
  saveStore();
}

export function listContainers() {
  return m.request('/api/v1/containers', {
    headers: {
      Authorization: `${store.token}`,
    },
  });
}

export function listImages() {
  return m.request('/api/v1/images');
}

export function currentUser() {
  return store.user;
}

export function startContainer(id) {
  return m.request('/api/v1/containers/start', {
    method: 'POST',
    headers: {
      Authorization: `${store.token}`,
    },
    params: { id },
  });
}

export function stopContainer(id) {
  return m.request(`/api/v1/containers/stop`, {
    method: 'POST',
    headers: {
      Authorization: `${store.token}`,
    },
    params: { id },
  });
}

export function destroyContainer(id) {
  return m.request(`/api/v1/containers/destroy`, {
    method: 'POST',
    headers: {
      Authorization: `${store.token}`,
    },
    params: { id },
  });
}

export function createContainer(image, customName) {
  return m.request(`/api/v1/containers`, {
    method: 'POST',
    headers: {
      Authorization: `${store.token}`,
    },
    body: { image, customName },
  });
}

export function listKeys() {
  return m.request('/api/v1/containers/keys', {
    headers: {
      Authorization: `${store.token}`,
    },
  });
}

export function generateKey(containerId) {
  return m.request('/api/v1/containers/keys', {
    method: 'POST',
    headers: {
      Authorization: `${store.token}`,
    },
    params: { id: containerId },
  });
}
