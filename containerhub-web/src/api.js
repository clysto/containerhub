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
    saveStore();
    return response;
  } catch (error) {
    throw error;
  }
}

export function listContainers() {
  return m.request('/api/v1/containers', {
    headers: {
      Authorization: `${store.token}`,
    },
  });
}
