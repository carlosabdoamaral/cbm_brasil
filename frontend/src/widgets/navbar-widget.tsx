import { useState } from "react";
import { Menu } from "semantic-ui-react";
import { COLORS } from "../enums/enums";

export function NavbarWidget() {
  const [routes, _] = useState([
    { key: 0, text: "Menu", to: "/" },
    { key: 1, text: "Nova ocorrência", to: "/new-occurrence" },
    { key: 2, text: "Criar conta", to: "/sign-up" },
  ]);

  function goToRoute(path: string) {
    window.location.href = path;
  }

  function isActiveRoute(path: string): boolean {
    return window.location.pathname === path;
  }

  return (
    <Menu fluid widths={6}>
      {routes.map((r, _) => (
        <Menu.Item
          active={isActiveRoute(r.to)}
          key={r.key}
          style={{
            color: COLORS.defaultRed,
          }}
          onClick={() => {
            goToRoute(r.to);
          }}
        >
          {r.text}
        </Menu.Item>
      ))}
    </Menu>
  );
}
