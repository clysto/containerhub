import Layout from '../components/layout';

function DocumentationPage() {
  return {
    view() {
      return (
        <Layout>
          <div class="border-bottom">
            <div className="container py-5">
              <h1>Documentation</h1>
              <p className="lead mb-0">
                Learn how to use ContainerHub.
              </p>
            </div>
          </div>
          <div class="container py-4">
            <p>Coming soon...</p>
          </div>
        </Layout>
      );
    },
  };
}

export default DocumentationPage;
