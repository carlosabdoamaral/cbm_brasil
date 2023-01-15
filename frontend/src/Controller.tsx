import { BrowserRouter, Route, Routes } from "react-router-dom";
import { HomePage } from "./pages/home-page";
import { NewOcurrencePage } from "./pages/new-occurrence-page";

export function Controller() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/new-occurrence" element={<NewOcurrencePage />} />
      </Routes>
    </BrowserRouter>
  );
}
