import { useEffect, useState } from "react";
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
import Spacer from "../widgets/spacer-widget";
import { NavbarWidget } from "../widgets/navbar-widget";
import axios from "axios";
import { addErrorNotification } from "../utils/notifications";
import OccurrenceCardWidget from "../widgets/occurrence-card-widget";

export function HomePage() {
  const maxOccurrences = 6;

  const [occurrences, setOccurrences]: any[] = useState([]);
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
          {!!occurrences &&
            occurrences.map(
              (item: any, i: number) =>
                i <= maxOccurrences && (
                  <GridColumn>
                    <OccurrenceCardWidget {...item} key={i} />
                  </GridColumn>
                )
            )}
        </Grid>
      );
    };

    const renderOccurrencesListUnlimited = () => {
      return (
        <Grid columns={3}>
          {!!occurrences &&
            occurrences.map((item: any, i: number) => (
              <GridColumn>
                <OccurrenceCardWidget {...item} key={i} />
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

  async function fetchAllOccurrences() {
    const URL = `${process.env.REACT_APP_OCCURRENCE_BASE_URL}/all`;
    await axios
      .get(URL)
      .then(({ data }) => {
        let dataList: any[] = [];
        data.map((item: any, i: number) => {
          let model = {
            IdOccurrence: item.id_occurrence,
            Title: item.title,
            Description: item.description,
            BannerX64: item.banner_x64,
            IdFirefighter: item.id_firefighter,
            FirefighterFullname: item.firefighter_fullname,
            FirefighterEmail: item.firefighter_email,
            IdAuthor: item.id_author,
            AuthorFullname: item.author_fullname,
            AuthorEmail: item.author_email,
            CreatedAt: item.created_at,
            UpdatedAt: item.updated_at,
            AcceptedAt: item.accepted_at,
            CanceledAt: item.canceled_at,
            FinishedAt: item.finished_at,
          };

          dataList.push(model);
        });

        setOccurrences(dataList);
      })
      .catch(({ response }) => {
        addErrorNotification(response.data);
      })
      .finally(() => {
        //code
      });
  }

  useEffect(() => {
    fetchAllOccurrences();
  }, []);

  return (
    <div style={{ paddingBottom: "3em" }}>
      <NavbarWidget />

      <Spacer height={60} />
      <Container>
        {renderHeader()}
        <Divider />
      </Container>

      <Spacer height={40} />
      <Container>{renderOccurrencesSection()}</Container>
    </div>
  );
}
