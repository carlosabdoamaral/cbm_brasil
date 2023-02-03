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
import { COLORS } from "../enums/enums";
import { SignInInterface } from "../interfaces/interfaces";
import { doPost } from "../utils/http";
import {
  addErrorNotification,
  addSuccessNotification,
} from "../utils/notifications";
import { renderMockModeHeader } from "../widgets/mock-widget";
import Spacer from "../widgets/spacer-widget";

export function SignInPage() {
  const [CPF, setCPF] = useState("");
  const [password, setPassword] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [mockModeActivated, setMockModeActivated] = useState(false);

  function handleOnCPFChange(cpf: string): string {
    cpf = cpf.replace(/[^\d]/g, "");

    return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
  }

  function someFieldIsEmpty() {
    return !CPF || !password;
  }

  async function handleSignInRequest() {
    if (someFieldIsEmpty()) {
      addErrorNotification("Todos os campos deve estar preenchidos");
      return;
    }

    const body: SignInInterface = {
      Cpf: CPF,
      Password: password,
    };

    setIsLoading(true);
    await doPost(
      process.env.REACT_APP_BACKEND_SIGN_IN_PATH || "/v1/auth/sign-in",
      body
    )
      .then((res) => {
        addSuccessNotification("Bem vindo novamente!");
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setIsLoading(false);
      });
  }

  function handleMockedSignInRequest() {
    if (someFieldIsEmpty()) {
      addErrorNotification("Todos os campos deve estar preenchidos");
      return;
    }

    const body: SignInInterface = {
      Cpf: CPF,
      Password: password,
    };

    setIsLoading(true);

    setTimeout(() => {
      addSuccessNotification(`CPF: ${body.Cpf} Pass: ${body.Password}`);
      setIsLoading(false);

      setTimeout(() => {
        window.location.href = "/";
      }, 1000);
    }, 2000);
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
          textAlign={"center"}
        >
          Olá!
        </Header>
      </>
    );
  }

  function renderForm() {
    return (
      <Form>
        <Grid columns={1}>
          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Seu CPF"}
              value={CPF}
              maxLength={11}
              onChange={(_, data) => setCPF(handleOnCPFChange(data.value))}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Input
              fluid
              placeholder={"Sua senha"}
              type={"password"}
              value={password}
              onChange={(_, data) => setPassword(data.value)}
            />
          </Grid.Column>

          <Grid.Column>
            <Form.Button
              fluid
              style={{ background: COLORS.defaultRed, color: "#FFF" }}
              content={"Entrar"}
              onClick={() => {
                mockModeActivated
                  ? handleMockedSignInRequest()
                  : handleSignInRequest();
              }}
            />
          </Grid.Column>
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
          width: "40vw",
          justifyContent: "center",
        }}
      >
        {mockModeActivated && renderMockModeHeader()}

        {renderHeader()}

        <Spacer height={50} />

        {renderForm()}

        <Header>
          <Header.Subheader>
            Não possui conta ainda?{" "}
            <Link to={"/sign-up"}>Crie uma agora mesmo</Link>
          </Header.Subheader>
        </Header>

        <Spacer height={50} />
      </Container>
    </>
  );
}
