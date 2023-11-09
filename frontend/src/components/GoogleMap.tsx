import React, { useEffect, useState } from "react";
import {
  GoogleMap,
  DirectionsRenderer,
  useLoadScript,
  Libraries,
} from "@react-google-maps/api";

const mapContainerStyle = {
  width: "100%",
  height: "600px",
};

const libraries: Libraries = ["places"];

function DirectionsDisplay({
  origin,
  destination,
  handleSetRoutes,
  directions,
  ...props
}: {
  origin: string;
  destination: string;
  handleSetRoutes: (route: google.maps.DirectionsResult) => void;
  directions: google.maps.DirectionsResult | null;
}) {
  const apiKey = process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY || "";
  const { isLoaded, loadError } = useLoadScript({
    googleMapsApiKey: apiKey,
    libraries,
  });
  const [center, setCenter] = useState<google.maps.LatLngLiteral>();
  useEffect(() => {
    if (!isLoaded) return;

    if (origin != "" && destination != "") {
      if (!origin.includes("station")) {
        origin = origin + " station";
      }
      if (!destination.includes("station")) {
        destination = destination + " station";
      }

      const directionsService = new google.maps.DirectionsService();

      directionsService.route(
        {
          origin,
          destination,
          travelMode: google.maps.TravelMode.TRANSIT,
        },
        (result, status) => {
          if (status === google.maps.DirectionsStatus.OK && result) {
            handleSetRoutes(result);
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

  if (loadError) return <div>Failed to load Google Maps.</div>;
  if (!isLoaded) return <div>Loading Google Maps...</div>;

  return (
    <div {...props}>
      <GoogleMap
        mapContainerStyle={mapContainerStyle}
        zoom={14}
        center={center}
      >
        {directions && <DirectionsRenderer directions={directions} />}
      </GoogleMap>
    </div>
  );
}

export default DirectionsDisplay;
