import React from "react";
import ReactDOM from "react-dom/client";
import reportWebVitals from "./reportWebVitals";

import "semantic-ui-css/semantic.min.css";
import { Controller } from "./controller";
import { COLORS } from "./enums/enums";

import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <ToastContainer
      position="top-right"
      autoClose={5000}
      hideProgressBar={false}
      newestOnTop={false}
      closeOnClick
      rtl={false}
      pauseOnFocusLoss
      draggable
      pauseOnHover
      theme="light"
    />

    <div style={{ height: "5px", background: COLORS.defaultRed }} />
    <Controller />

    <ToastContainer />
  </React.StrictMode>
);

reportWebVitals();
