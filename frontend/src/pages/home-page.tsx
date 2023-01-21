import { useState } from "react";
import { Link } from "react-router-dom";
import {
  Button,
  Container,
  Divider,
  Grid,
  GridColumn,
  Header,
  Image,
  Segment,
} from "semantic-ui-react";
import { COLORS, MOCK_TEMPLATES } from "../enums/enums";
import { GenerateMock } from "../mocks/gen";
import OccurrenceCardWidget from "../widgets/occurrence-card-widget";
import Spacer from "../widgets/spacer-widget";
import firefighters from "../assets/images/firefighters.jpeg";
import { NavbarWidget } from "../widgets/navbar-widget";

export function HomePage() {
  const maxOccurrences = 6;

  const [occurrences, setOccurrences] = useState(
    GenerateMock(MOCK_TEMPLATES.OCCURRENCES)
  );
  const [isShowingAllOccurrences, setIsShowingAllOccurrences] = useState(false);

  function renderHeader() {
    return (
      <>
        <Header
          as="h1"
          style={{ fontSize: 35, fontFamily: "poppins", fontWeight: 500 }}
        >
          Menu
          <Header.Subheader style={{ fontFamily: "poppins", fontWeight: 200 }}>
            Acompanhe todas ocorrências próximas de você
          </Header.Subheader>
        </Header>
      </>
    );
  }

  function renderOccurrencesSection() {
    const renderSectionTitle = () => {
      return (
        <Container
          style={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "top",
          }}
        >
          <Header
            as="h2"
            style={{
              fontFamily: "poppins",
              fontWeight: 500,
            }}
          >
            Ocorrências
            <Header.Subheader
              style={{
                fontFamily: "poppins",
                fontWeight: 300,
              }}
            >
              Mais recentes
            </Header.Subheader>
          </Header>
          <Link
            to={"/new-occurrence"}
            style={{
              background: COLORS.defaultRed,
              height: "fit-content",
              padding: ".5em",
              borderRadius: 5,
              color: "white",
              marginTop: 10,
            }}
          >
            Criar ocorrência
          </Link>
        </Container>
      );
    };

    const renderOccurrencesListLimited = () => {
      return (
        <Grid columns={3}>
          {occurrences!.map(
            (occurrence, index) =>
              index < maxOccurrences && (
                <GridColumn>
                  <OccurrenceCardWidget {...occurrence} />
                </GridColumn>
              )
          )}
        </Grid>
      );
    };

    const renderOccurrencesListUnlimited = () => {
      return (
        <Grid columns={3}>
          {occurrences!.map((occurrence, index) => (
            <GridColumn>
              <OccurrenceCardWidget {...occurrence} />
            </GridColumn>
          ))}
        </Grid>
      );
    };

    return (
      <>
        {renderSectionTitle()}
        <Spacer height={20} />

        {isShowingAllOccurrences
          ? renderOccurrencesListUnlimited()
          : renderOccurrencesListLimited()}

        <Spacer height={20} />

        {isShowingAllOccurrences && (
          <Button
            fluid
            content={"Exibir as 6 mais recentes"}
            icon={"arrow up"}
            onClick={() => {
              setIsShowingAllOccurrences(!isShowingAllOccurrences);
            }}
          />
        )}

        {!isShowingAllOccurrences && occurrences!.length > 6 && (
          <Button
            fluid
            content={`Exibir todas as ${occurrences?.length} ocorrências`}
            icon={"arrow down"}
            onClick={() => {
              setIsShowingAllOccurrences(!isShowingAllOccurrences);
            }}
          />
        )}
      </>
    );
  }

  return (
    <div style={{ paddingBottom: "3em" }}>
      <NavbarWidget />

      <Spacer height={60} />
      <Container>{renderHeader()}</Container>

      <Container>
        <Divider />
      </Container>

      <Spacer height={40} />
      <Container>{renderOccurrencesSection()}</Container>
    </div>
  );
}
