import { Header } from "semantic-ui-react";
import Spacer from "./spacer-widget";

export function renderMockModeHeader() {
    return (
      <>
        <Header as={"h1"} color={"red"} textAlign="center">
          Mock mode activated
        </Header>
        <Spacer height={50} />
      </>
    );
  }