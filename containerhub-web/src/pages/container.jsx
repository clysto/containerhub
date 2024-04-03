import Layout from '../components/layout';
import ContainerCard from '../components/container-card';
import { listContainers } from '../api';

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
          <div className="container p-4">
            <h1 class="mb-4">My Containers</h1>
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
