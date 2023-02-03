import axios from "axios";

export async function doPost(
  path: string,
  body: object,
  config?: object | undefined
) {
  const url = `${process.env.REACT_APP_BACKEND_BASE_URL}${path}`;

  await axios
    .post(url, body, config)
    .then((res) => {
      console.log(res.data);
    })
    .catch((err) => {
      console.error(err);
    });
}

export async function doAwaitGet(url: string, config: object) {
  await axios
    .get(url, config)
    .then((res) => {
      return res;
    })
    .catch((err) => {
      return err;
    });
}
