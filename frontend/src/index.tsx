import React from "react";
import ReactDOM from "react-dom/client";
import reportWebVitals from "./reportWebVitals";

import "semantic-ui-css/semantic.min.css";
import "react-notifications-component/dist/theme.css";
import { ReactNotifications } from "react-notifications-component";
import { Controller } from "./controller";
import { COLORS } from "./enums/enums";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <ReactNotifications />
    <>
      <div style={{ height: "5px", background: COLORS.defaultRed }} />
      <Controller />
    </>
  </React.StrictMode>
);

reportWebVitals();
