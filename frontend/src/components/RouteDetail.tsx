import React from "react";
import { Paper, Typography, makeStyles, Theme } from "@mui/material";

const RouteDetails = ({ route }: { route: google.maps.DirectionsRoute }) => {
  const legs = route.legs;
  return (
    <Paper>
      <Typography variant="h6">乗換案内</Typography>
      <Typography>出発地点: {legs[0].start_address}</Typography>
      <Typography>目的地: {legs[0].end_address}</Typography>
      <Typography>距離: {legs[0].distance?.text}</Typography>
      <Typography>所要時間: {legs[0].duration?.text}</Typography>
      <Typography>出発時刻: {legs[0].departure_time?.text}</Typography>
      <Typography>到着時刻: {legs[0].arrival_time?.text}</Typography>
    </Paper>
  );
};

export default RouteDetails;
