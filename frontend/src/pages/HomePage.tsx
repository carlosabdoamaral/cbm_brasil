import { Container, Divider } from "semantic-ui-react";
import { TitleWidget } from "../widgets/TitleWidget";

export function HomePage() {
  return (
    <Container>
      <TitleWidget title={"Menu"} subtitle={""} />
      <Divider />
    </Container>
  );
}
