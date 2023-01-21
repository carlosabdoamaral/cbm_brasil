import { useState } from "react";
import { Link } from "react-router-dom";
import {
  Button,
  Container,
  Dimmer,
  Form,
  Grid,
  Header,
  Icon,
  Loader,
  Select,
} from "semantic-ui-react";
import {
  AddressInterface,
  NewOccurrenceRequestInterface,
} from "../interfaces/interfaces";
import {
  addressInterfaceIsEmpty,
  emtpyAddr,
  formatCEP,
  getCEPDetails,
} from "../utils/cep";
import { doPost } from "../utils/http";
import {
  addErrorNotification,
  addSuccessNotification,
  addWarningNotification,
} from "../utils/notifications";
import { GetOperationalSystem } from "../utils/operational-system";
import Spacer from "../widgets/spacer-widget";

export function NewOcurrencePage() {
  const categoryOptions = [
    { key: "FIRE", text: "Incêndio", value: "FIRE" },
    { key: "CAR_FIRE", text: "Carro em chamas", value: "CAR_FIRE" },
    { key: "ANIMALS", text: "Animal em perigo", value: "ANIMALS" },
  ];

  const [title, setTitle] = useState(String);
  const [image, setImage] = useState(String);
  const [CEP, setCEP] = useState(String);
  const [category, setCategory] = useState(Object);
  const [isLoading, setIsLoading] = useState(false);
  const [country, setCountry] = useState(String);
  const [state, setState] = useState(String);
  const [city, setCity] = useState(String);
  const [neighborhood, setNeighborhood] = useState(String);
  const [street, setStreet] = useState(String);
  const [placeNumber, setPlaceNumber] = useState(String);
  const [placeNotes, setPlaceNotes] = useState(String);

  function allRequiredFieldsAreFilled(): boolean {
    return !!title && !!image && !!CEP && !!category;
  }

  function updateAddressInfos(infos: AddressInterface) {
    setCountry(infos.Country);
    setState(infos.State);
    setCity(infos.City);
    setNeighborhood(infos.Neighborhood);
    setStreet(infos.Street);
  }

  async function handleSubmit() {
    if (allRequiredFieldsAreFilled()) {
      setIsLoading(true);

      let addr = await getCEPDetails(CEP);
      if (addressInterfaceIsEmpty(addr)) {
        return;
      }

      const body: NewOccurrenceRequestInterface = {
        AccountId: 1,
        CreatedAt: new Date(),
        System: GetOperationalSystem(),
        Title: title,
        Image: image,
        Location: addr,
      };

      doPost("/api/", body)
        .then((res) => {
          addSuccessNotification("Enviado com sucesso");
          handleClearValues();
        })
        .catch((err) => {
          addErrorNotification(err);
        })
        .finally(() => {
          setIsLoading(false);
        });
    } else {
      addErrorNotification("Todos os campos devem estar preenchidos");
    }
  }

  function handleClearValues(mustNotify: boolean = false) {
    setTitle("");
    setImage("");
    setCEP("");
    setCountry("");
    setState("");
    setCity("");
    setNeighborhood("");
    setStreet("");
    setPlaceNumber("");
    setPlaceNotes("");
    setCategory({ key: "", text: "", value: "" });

    if (mustNotify) {
      addSuccessNotification("Todos os campos foram limpos");
    }
  }

  function renderHeader() {
    return (
      <>
        <Link to={"/"}>
          <Icon name="arrow left" color="black" />
        </Link>
        <Header
          as="h1"
          style={{ fontSize: 35, fontFamily: "poppins", fontWeight: 500 }}
        >
          Crie sua ocorrência <br /> agora mesmo 🧯
        </Header>
      </>
    );
  }

  function renderForm() {
    return (
      <Form>
        <Grid columns={2}>
          <Grid.Column>
            <Form.Input
              fluid
              label="Título"
              placeholder="Título"
              value={title}
              onChange={(_, data) => {
                setTitle(data.value);
              }}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Select
              label="Categoria"
              onChange={(_, data) => {
                const category = categoryOptions.filter(
                  (category) => category.value === data.value
                );

                setCategory(category);
              }}
              options={categoryOptions}
              placeholder="Categoria"
              value={category.value}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              label={"CEP"}
              placeholder="CEP"
              value={CEP}
              maxLength={8}
              onChange={async (_, data) => {
                updateAddressInfos(emtpyAddr);
                setCEP(formatCEP(data.value));

                if (data.value.length >= 8) {
                  let addr = await getCEPDetails(data.value);
                  if (addressInterfaceIsEmpty(addr)) {
                    return;
                  }

                  updateAddressInfos(addr);
                }
              }}
            />
            <small>
              {!!CEP &&
              CEP.length > 8 &&
              !!country &&
              !!state &&
              !!city &&
              !!street
                ? `${country}, ${state}, ${city}, ${street}`
                : ""}
            </small>
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              type="file"
              fluid
              label="Imagem"
              placeholder="Imagem"
              onChange={(_, data) => {
                setImage(data.value);
              }}
            />
          </Grid.Column>
        </Grid>

        <Spacer height={20} />

        <Spacer height={10} />

        <Grid columns={2}>
          <Grid.Row>
            <Grid.Column>
              <Button
                fluid
                content="Limpar valores"
                color="red"
                icon={"close"}
                onClick={() => {
                  handleClearValues(true);
                }}
              />
            </Grid.Column>

            <Grid.Column>
              <Button
                fluid
                content="Enviar"
                secondary
                icon={"send"}
                onClick={() => {
                  handleSubmit();
                }}
              />
            </Grid.Column>
          </Grid.Row>
        </Grid>
      </Form>
    );
  }

  return (
    <>
      <Dimmer active={isLoading}>
        <Loader active={isLoading} />
      </Dimmer>
      <Container
        style={{
          display: "flex",
          flexDirection: "column",
          height: "100vh",
          justifyContent: "center",
        }}
      >
        <Container>
          {renderHeader()}
          <Container style={{ height: "30px" }} />
          {renderForm()}
        </Container>
      </Container>
    </>
  );
}
