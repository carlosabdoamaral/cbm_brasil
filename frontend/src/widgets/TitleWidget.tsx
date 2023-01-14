import { Container, Header } from "semantic-ui-react";
import { TitleWidgetProps } from "../interfaces/TitleWidgetInterfaces";

export function TitleWidget(props: TitleWidgetProps) {
  return (
    <Container style={{ marginTop: "3em" }}>
      <Header as={"h2"}>{props.title}</Header>
      <Header.Subheader>{props.subtitle}</Header.Subheader>
    </Container>
  );
}
