import Snackbar from "@mui/material/Snackbar";
import React from "react";
import MuiAlert, { AlertProps, AlertColor } from "@mui/material/Alert";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
  props,
  ref
) {
  return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

type CustomSnackBarProps = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  severity: AlertColor;
  text: string;
};

export enum Severity {
  ERROR = "error",
  WARNING = "warning",
  INFO = "info",
  SUCCESS = "success",
}

export default function CustomSnackBar({
  open,
  setOpen,
  severity,
  text,
}: CustomSnackBarProps) {
  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }

    setOpen(false);
  };
  return (
    <section>
      <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
        <Alert severity={severity} sx={{ width: "100%" }} onClose={handleClose}>
          {text}
        </Alert>
      </Snackbar>
    </section>
  );
}
