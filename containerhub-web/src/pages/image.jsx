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
          <div className="container p-4">
            <h1 class="mb-4">All Images</h1>
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
