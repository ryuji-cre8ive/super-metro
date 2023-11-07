import * as React from "react";
import Box from "@mui/material/Box";
import Input from "@mui/material/Input";
import InputAdornment from "@mui/material/InputAdornment";
import FormControl from "@mui/material/FormControl";
import TextField from "@mui/material/TextField";
import AccountCircle from "@mui/icons-material/AccountCircle";
import DirectionsTransitIcon from "@mui/icons-material/DirectionsTransit";

interface InputWithIconProps {
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  label: InputLabel;
}

export enum InputLabel {
  DEPARTURE = "DEPARTURE",
  DESTINATION = "DESTINATION",
}

export default function InputWithIcon({ onChange, label }: InputWithIconProps) {
  return (
    <Box sx={{ "& > :not(style)": { m: 1 } }}>
      <TextField
        id="input-with-icon-textfield"
        label={label}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <DirectionsTransitIcon />
            </InputAdornment>
          ),
        }}
        onChange={onChange}
        variant="standard"
      />
    </Box>
  );
}
