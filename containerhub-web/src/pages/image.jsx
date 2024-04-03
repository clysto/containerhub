import { listImages } from '../api';
import ImageCard from '../components/image-card';
import Layout from '../components/layout';

function ImagePage() {
  let images = [];
  return {
    oninit() {
      listImages().then((data) => {
        images = data;
      });
    },
    view() {
      return (
        <Layout>
          <div class="border-bottom">
            <div className="container py-5">
              <h1>All Images</h1>
              <p className="lead mb-0">Browse all available images.</p>
            </div>
          </div>
          <div className="container py-4">
            <div className="row row-cols-1 row-cols-lg-2 row-cols-xl-3 g-4">
              {images.map((image) => (
                <div class="col">
                  <ImageCard image={image} />
                </div>
              ))}
            </div>
          </div>
        </Layout>
      );
    },
  };
}

export default ImagePage;
