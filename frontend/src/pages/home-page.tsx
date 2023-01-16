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
        {isShowingAllOccurrences ? (
          <Button
            fluid
            content={"Exibir as 6 mais recentes"}
            icon={"arrow up"}
            onClick={() => {
              setIsShowingAllOccurrences(!isShowingAllOccurrences);
            }}
          />
        ) : (
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

  function renderProjectObjectiveSection() {
    return (
      <div style={{ background: COLORS.defaultRed, padding: "3em" }}>
        <Container>
          <Header
            as={"h2"}
            content="Objetivo do projeto"
            style={{ color: "#FFF" }}
          />
          <p style={{ color: "#FFF", fontFamily: "poppins", fontWeight: 200 }}>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Ut a dolore
            facere laboriosam ducimus possimus assumenda vitae eius, provident
            in quidem earum quis. Veritatis quas ex earum, natus omnis adipisci.
            Assumenda, fugiat. Laboriosam explicabo cupiditate, beatae
            voluptatem atque dolor facilis. Distinctio inventore facere nam
            minima fugiat eius natus non itaque? Modi vero repudiandae
            perferendis autem nemo sunt atque, officiis eos. Lorem ipsum dolor
            sit amet consectetur adipisicing elit. Ut a dolore facere laboriosam
            ducimus possimus assumenda vitae eius, provident in quidem earum
            quis. Veritatis quas ex earum, natus omnis adipisci. Assumenda,
            fugiat. Laboriosam explicabo cupiditate, beatae voluptatem atque
            dolor facilis. Distinctio inventore facere nam minima fugiat eius
            natus non itaque? Modi vero repudiandae perferendis autem nemo sunt
            atque, officiis eos.
          </p>
        </Container>
      </div>
    );
  }

  function renderHowToUseSection() {
    return (
      <Container>
        <Header as={"h2"} content="Como usar a plataforma" />
        <p style={{ fontFamily: "poppins", fontWeight: 200 }}>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Ut a dolore
          facere laboriosam ducimus possimus assumenda vitae eius, provident in
          quidem earum quis. Veritatis quas ex earum, natus omnis adipisci.
          Assumenda, fugiat. Laboriosam explicabo cupiditate, beatae voluptatem
          atque dolor facilis. Distinctio inventore facere nam minima fugiat
          eius natus non itaque? Modi vero repudiandae perferendis autem nemo
          sunt atque, officiis eos. Lorem ipsum dolor sit amet consectetur
          adipisicing elit. Ut a dolore facere laboriosam ducimus possimus
          assumenda vitae eius, provident in quidem earum quis. Veritatis quas
          ex earum, natus omnis adipisci. Assumenda, fugiat. Laboriosam
          explicabo cupiditate, beatae voluptatem atque dolor facilis.
          Distinctio inventore facere nam minima fugiat eius natus non itaque?
          Modi vero repudiandae perferendis autem nemo sunt atque, officiis eos.
        </p>
      </Container>
    );
  }

  return (
    <div style={{ paddingBottom: "3em" }}>
      <Spacer height={60} />
      <Container>{renderHeader()}</Container>

      <Spacer height={40} />
      <Container>
        <Divider />
      </Container>

      <Spacer height={40} />
      {renderHowToUseSection()}
      
      <Spacer height={40} />
      <Container>
        <Divider />
      </Container>

      <Spacer height={40} />
      <Container>{renderOccurrencesSection()}</Container>

      <Spacer height={40} />
      {renderProjectObjectiveSection()}

      <Spacer height={40} />
      <Container>
        <Divider />
      </Container>
    </div>
  );
}
