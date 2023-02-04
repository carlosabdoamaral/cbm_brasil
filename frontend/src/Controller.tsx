import { BrowserRouter, Route, Routes } from "react-router-dom";
import { HomePage } from "./pages/home-page";
import { NewOcurrencePage } from "./pages/new-occurrence-page";
import { SignInPage } from "./pages/sign-in-page";
import { SignUpPage } from "./pages/sign-up-page";

export function Controller() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/new-occurrence" element={<NewOcurrencePage />} />
        <Route path="/sign-in" element={<SignInPage />} />
        <Route path="/sign-up" element={<SignUpPage />} />
      </Routes>
    </BrowserRouter>
  );
}
