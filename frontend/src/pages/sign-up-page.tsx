import { useState } from "react";
import { Link } from "react-router-dom";
import {
  Container,
  Dimmer,
  Form,
  Grid,
  Header,
  Icon,
  Loader,
} from "semantic-ui-react";
import { COLORS, SIGN_UP_FORMS } from "../enums/enums";
import { AddressInterface } from "../interfaces/interfaces";
import { getCEPDetails } from "../utils/cep";
import {
  addErrorNotification,
  addSuccessNotification,
} from "../utils/notifications";
import { renderMockModeHeader } from "../widgets/mock-widget";
import Spacer from "../widgets/spacer-widget";

export function SignUpPage() {
  const [CPF, setCPF] = useState(String);
  const [fullName, setFullName] = useState(String);
  const [age, setAge] = useState(String);
  const [password, setPassword] = useState(String);

  const [CEP, setCEP] = useState(String);
  const [country, setCountry] = useState(String);
  const [state, setState] = useState(String);
  const [city, setCity] = useState(String);
  const [neighborhood, setNeighborhood] = useState(String);
  const [street, setStreet] = useState(String);
  const [placeNumber, setPlaceNumber] = useState(String);
  const [placeNotes, setPlaceNotes] = useState(String);

  const [isLoading, setIsLoading] = useState(false);
  const [mockModeActivated, setMockModeActivated] = useState(false);
  const [activeForm, setActiveForm] = useState(SIGN_UP_FORMS.personalInfos);
  const formsOptions = [
    SIGN_UP_FORMS.personalInfos,
    SIGN_UP_FORMS.addressInfos,
  ];

  function handleOnCPFChange(cpf: string) {
    cpf = cpf.replace(/[^\d]/g, "");
    setCPF(cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4"));
    return;
  }

  function updateAddressInfos(infos: AddressInterface) {
    setCountry(infos.Country);
    setState(infos.State);
    setCity(infos.City);
    setNeighborhood(infos.Neighborhood);
    setStreet(infos.Street);
  }

  async function handleOnCEPChange(inputValue: string) {
    var re = /^([\d]{5})\.*-*([\d]{3})/; // Pode usar ? no lugar do *
    if (re.test(inputValue)) {
      setCEP(inputValue.replace(re, "$1-$2"));

      let addr = await getCEPDetails(inputValue);
      if (addr.Country === "EMPTY") {
        addErrorNotification(
          "Não foi possível encontrar infomações do CEP fornecido"
        );

        return;
      }

      updateAddressInfos(addr);
    } else {
      setCEP(inputValue);
    }
  }

  function handleClearForm() {
    setCPF("");
    setFullName("");
    setAge("");
    setPassword("");
    setCEP("");
    setCountry("");
    setState("");
    setCity("");
    setNeighborhood("");
    setStreet("");
    setPlaceNumber("");
    setPlaceNotes("");
  }

  function renderPersonalInfosHeader() {
    return (
      <>
        <Header
          as="h1"
          style={{ fontSize: 35, fontFamily: "poppins", fontWeight: 500 }}
        >
          Olá! <br />
          Crie sua conta 🚀
        </Header>
      </>
    );
  }

  function renderAddressInfosHeader() {
    return (
      <>
        <Icon
          name="arrow left"
          color="black"
          onClick={() => {
            setActiveForm(formsOptions[0]);
          }}
        />
        <Header
          as="h1"
          style={{ fontSize: 35, fontFamily: "poppins", fontWeight: 500 }}
        >
          Estamos quase <br />
          Finalizando! 🎉
        </Header>
      </>
    );
  }

  function renderHeader() {
    return (
      <>
        {activeForm === SIGN_UP_FORMS.personalInfos &&
          renderPersonalInfosHeader()}

        {activeForm === SIGN_UP_FORMS.addressInfos &&
          renderAddressInfosHeader()}
      </>
    );
  }

  function renderPersonalInfosForm() {
    return (
      <Form>
        <Grid columns={2}>
          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Nome completo"}
              type={"text"}
              value={fullName}
              onChange={(_, data) => setFullName(data.value)}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Seu CPF"}
              value={CPF}
              type={"text"}
              maxLength={14}
              onChange={(_, data) => handleOnCPFChange(data.value)}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Idade"}
              type={"number"}
              value={age}
              onChange={(_, data) => setAge(data.value)}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Sua senha"}
              type={"text"}
              value={password}
              onChange={(_, data) => setPassword(data.value)}
            />
          </Grid.Column>

          <Grid.Column>
            <Link to={"/"}>
              <Form.Button fluid content={"Voltar ao menu"} />
            </Link>
          </Grid.Column>

          <Grid.Column>
            <Form.Button
              fluid
              style={{ background: COLORS.defaultRed, color: "#FFF" }}
              content={"Próxima etapa"}
              onClick={() => {
                setActiveForm(formsOptions[1]);
              }}
            />
          </Grid.Column>
        </Grid>
      </Form>
    );
  }

  function renderAddressInfosForm() {
    return (
      <Form>
        <Grid columns={2}>
          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"CEP"}
              type={"text"}
              value={CEP}
              maxLength={9}
              onChange={(_, data) => {
                handleOnCEPChange(data.value);
              }}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"País"}
              readOnly
              type={"text"}
              value={country}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Estado"}
              readOnly
              type={"text"}
              value={state}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Cidade"}
              readOnly
              type={"text"}
              value={city}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Bairro"}
              readOnly
              type={"text"}
              value={neighborhood}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Rua"}
              readOnly
              type={"text"}
              value={street}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Número"}
              type={"number"}
              value={placeNumber}
              onChange={(_, data) => {
                setPlaceNumber(data.value);
              }}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Complemento"}
              type={"text"}
              value={placeNotes}
              onChange={(_, data) => {
                setPlaceNotes(data.value);
              }}
            />
          </Grid.Column>
        </Grid>

        <Spacer height={25} />
        <Form.Button
          fluid
          style={{ background: COLORS.defaultRed, color: "#FFF" }}
          content={"Finalizar cadastro"}
          onClick={() => {
            addSuccessNotification(
              "Bem vindo! Sinta-se em casa e lembre-se, estamos aqui para ajudar você sempre!"
            );
          }}
        />
      </Form>
    );
  }

  function renderForm() {
    return (
      <>
        {activeForm === SIGN_UP_FORMS.personalInfos && (
          <>
            <Header>
              <Header.Subheader>Informações pessoais</Header.Subheader>
            </Header>
            {renderPersonalInfosForm()}
          </>
        )}

        {activeForm === SIGN_UP_FORMS.addressInfos && (
          <>
            <>
              <Header>
                <Header.Subheader>Informações do endereço</Header.Subheader>
              </Header>
              {renderAddressInfosForm()}
            </>
          </>
        )}
      </>
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
        {mockModeActivated && renderMockModeHeader()}
        {renderHeader()}
        {renderForm()}
      </Container>
    </>
  );
}
