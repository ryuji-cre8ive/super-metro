import React from "react";
import { Card, Typography, makeStyles, Theme } from "@mui/material";

const RouteDetails = ({ route }: { route: google.maps.DirectionsRoute }) => {
  const legs = route.legs;
  return (
    <Card style={{ width: "50%", display: "flex", padding: "30px" }}>
      <section style={{ flexGrow: "7" }}>
        <Typography>
          {legs[0].departure_time?.text} → {legs[0].arrival_time?.text}
        </Typography>
        <Typography>OneWay</Typography>
        <Typography>Steps: {legs[0].steps.length}</Typography>
      </section>
      <div
        style={{
          margin: "0 auto",
          display: "flex",
          alignItems: "center",
          flexGrow: "3",
        }}
      >
        <Typography>{legs[0].duration?.text}</Typography>
      </div>

      {/* <Typography>出発地点: {legs[0].start_address}</Typography>
      <Typography>目的地: {legs[0].end_address}</Typography>
      <Typography>距離: {legs[0].distance?.text}</Typography>
      <Typography>所要時間: {legs[0].duration?.text}</Typography>
      <Typography>出発時刻: {legs[0].departure_time?.text}</Typography>
      <Typography>到着時刻: {legs[0].arrival_time?.text}</Typography> */}
    </Card>
  );
};

export default RouteDetails;
