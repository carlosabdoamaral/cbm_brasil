import { Container, Divider, Header } from "semantic-ui-react";
import { TitleWidgetProps } from "../interfaces/TitleWidgetInterfaces";

export function TitleWidget(props: TitleWidgetProps) {
  return (
    <div style={{ margin: "2em 0" }}>
      <Header as={"h2"}>{props.title}</Header>
      <Header.Subheader>{props.subtitle}</Header.Subheader>
    </div>
  );
}
