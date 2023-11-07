import React, { useEffect, useState } from "react";
import {
  GoogleMap,
  DirectionsRenderer,
  useLoadScript,
  Libraries,
} from "@react-google-maps/api";

import RouteDetails from "./RouteDetail";

const mapContainerStyle = {
  width: "800px",
  height: "600px",
};

const libraries: Libraries = ["places"];

function DirectionsDisplay({
  origin,
  destination,
}: {
  origin: string;
  destination: string;
}) {
  const apiKey = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY || "";
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: apiKey,
    libraries,
  });

  const [directions, setDirections] =
    useState<google.maps.DirectionsResult | null>(null);
  const [center, setCenter] = useState<google.maps.LatLngLiteral>();

  useEffect(() => {
    if (!isLoaded) return;

    if (origin && destination) {
      console.log("origingin");
      const directionsService = new google.maps.DirectionsService();

      directionsService.route(
        {
          origin,
          destination,
          travelMode: google.maps.TravelMode.TRANSIT,
        },
        (result, status) => {
          if (status === google.maps.DirectionsStatus.OK) {
            setDirections(result);
            console.log(result);
          } else {
            alert("something error occurred, please check your query");
            console.error(`Directions request failed due to ${status}`);
          }
        }
      );
    }

    navigator.geolocation.getCurrentPosition((position) => {
      setCenter({
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      });
    });
  }, [origin, destination, isLoaded]);

  if (loadError) return <div>Google Mapsの読み込みに失敗しました。</div>;
  if (!isLoaded) return <div>Google Mapsを読み込んでいます...</div>;

  return (
    <div>
      <GoogleMap
        mapContainerStyle={mapContainerStyle}
        zoom={14}
        center={center}
      >
        {directions && <DirectionsRenderer directions={directions} />}
      </GoogleMap>
      <div>
        <h1>Routes</h1>
      </div>
      {directions &&
        directions.routes.map((route, index) => (
          <RouteDetails route={route} key={index} />
        ))}
    </div>
  );
}

export default DirectionsDisplay;
