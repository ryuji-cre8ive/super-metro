"use client";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { jwtDecode } from "jwt-decode";

import Map from "@/components/GoogleMap";

export default function Page({ params }: { params: { id: string } }) {
  const router = useRouter();
  const [origin, setOrigin] = useState<string>("");
  const [destination, setDestination] = useState<string>("");
  const [originInput, setOriginInput] = useState<string>("");
  const [destinationInput, setDestinationInput] = useState<string>("");

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
  return (
    <div style={{ display: "flex" }}>
      <section>
        <Map origin={origin} destination={destination} />
      </section>
      <section>
        <input type="text" onChange={handleChangeOrigin} />
        {origin}
        <input type="text" onChange={handleChangeDestination} />
        {destination}
        <button onClick={handleClickButton}>Search</button>
      </section>
    </div>
  );
}
