"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { jwtDecode } from "jwt-decode";

import Map from "@/components/GoogleMap";
import InputWithIcon, { InputLabel } from "@/components/InputWithIcon";
import { Button } from "@mui/material";
import RouteDetails from "@/components/RouteDetail";

export default function Page({ params }: { params: { id: string } }) {
  const router = useRouter();
  const [origin, setOrigin] = useState<string>("");
  const [destination, setDestination] = useState<string>("");
  const [originInput, setOriginInput] = useState<string>("");
  const [destinationInput, setDestinationInput] = useState<string>("");

  const [directions, setDirections] =
    useState<google.maps.DirectionsResult | null>(null);

  const handleChangeOrigin = (e: React.ChangeEvent<HTMLInputElement>) => {
    setOriginInput(e.target.value);
  };

  const handleChangeDestination = (e: React.ChangeEvent<HTMLInputElement>) => {
    setDestinationInput(e.target.value);
  };

  const handleClickButton = () => {
    setOrigin(originInput);
    setDestination(destinationInput);
  };

  const handleSetRoutes = (route: google.maps.DirectionsResult) => {
    setDirections(route);
  };
  return (
    <div>
      <section
        style={{
          display: "flex",
          textAlign: "center",
          justifyContent: "center",
        }}
      >
        <InputWithIcon
          onChange={handleChangeOrigin}
          label={InputLabel.DEPARTURE}
        />
        <InputWithIcon
          onChange={handleChangeDestination}
          label={InputLabel.DESTINATION}
        />
        <Button onClick={handleClickButton}>Search!!</Button>
      </section>
      <section style={{ display: "flex" }}>
        <div style={{ flex: "7", margin: "10px" }}>
          <Map
            origin={origin}
            destination={destination}
            handleSetRoutes={handleSetRoutes}
            directions={directions}
          />
        </div>

        <section style={{ flex: "3", margin: "10px" }}>
          <h1>Routes</h1>
          {directions &&
            directions.routes.map((route, index) => (
              <RouteDetails route={route} key={index} />
            ))}
        </section>
      </section>
    </div>
  );
}
