import React from "react";
import ReactDOM from "react-dom/client";
import reportWebVitals from "./reportWebVitals";

import "semantic-ui-css/semantic.min.css";
import "react-notifications-component/dist/theme.css";
import { ReactNotifications } from "react-notifications-component";
import { Controller } from "./Controller";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <ReactNotifications />
    <div>
      <div style={{ height: "5px", background: "red" }} />
      <Controller />
    </div>
  </React.StrictMode>
);

reportWebVitals();
