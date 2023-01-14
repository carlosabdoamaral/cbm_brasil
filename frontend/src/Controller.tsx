import { BrowserRouter, Route, Routes } from "react-router-dom";
import { HomePage } from "./pages/HomePage";
import { NewOcurrencePage } from "./pages/NewOccurrencePage";

export function Controller() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/new-ocurrence" element={<NewOcurrencePage />} />
      </Routes>
    </BrowserRouter>
  );
}
