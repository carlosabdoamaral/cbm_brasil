import { Container } from "semantic-ui-react";

interface SpacerInterface {
  height: number;
}

export default function Spacer(distance: SpacerInterface) {
  return <Container style={{ height: distance.height }} />;
}
