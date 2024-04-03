import Layout from '../components/layout';
import ContainerCard from '../components/container-card';
import { listContainers } from '../api';
import { Link } from 'mithril/route';

function ContainerPage() {
  let containers = [];

  function handleCardChange() {
    listContainers().then((data) => {
      containers = data;
      m.redraw();
    });
  }

  return {
    oninit() {
      listContainers().then((data) => {
        containers = data;
      });
    },
    view() {
      return (
        <Layout>
          <div class="border-bottom">
            <div className="container py-5">
              <h1>My Containers</h1>
              <p className="lead mb-0">
                Manage your containers. <Link href="/images">Create a new container</Link>.
              </p>
            </div>
          </div>
          <div className="container py-4">
            {containers.length === 0 && (
              <div class="alert alert-light" role="alert">
                You don't have any containers yet.
              </div>
            )}
            <div className="row row-cols-1 row-cols-lg-2 row-cols-xl-3 g-4">
              {containers.map((container) => (
                <div class="col" key={container.Id}>
                  <ContainerCard container={container} onchange={handleCardChange} />
                </div>
              ))}
            </div>
          </div>
        </Layout>
      );
    },
  };
}

export default ContainerPage;
