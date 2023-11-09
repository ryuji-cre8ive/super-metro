"use client";
import React, { useState } from "react";
import { TextField, Typography } from "@mui/material";
import LoadingButton from "@mui/lab/LoadingButton";
import SendIcon from "@mui/icons-material/Send";
import Image from "next/image";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
import EWalletPic from "../../../../../public/e-wallet.png";
import { SnackBarTextErr, SnackBarTextSuccess } from "./text";
import { useAuth } from "@/app/userContext";
import axios from "@/api/axiosConfig";
import { AxiosResponse } from "axios";

const TopUp = () => {
  const { user, topUp } = useAuth();
  const [amount, setAmount] = useState<number | null>(null);
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);

  const isNumber = (value: any) => {
    var pattern = /^[0-9]*$/;
    return pattern.test(value);
  };

  const handleAmount = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    if (!isNumber(e.target.value)) {
      return setErrOpen(true);
    }
    setErrOpen(false);
    setAmount(Number(e.target.value));
  };

  const handleTopUp = () => {
    if (!amount) return;
    if (!isNumber(amount)) {
      return setErrOpen(true);
    }
    if (!user) return;
    setLoading(true);
    try {
      axios
        .post("/top-up", { valance: Number(amount), id: user.id })
        .then((res: AxiosResponse) => {
          if (res.status !== 200) return;
          topUp(user, Number(amount));
          window.localStorage.setItem("session_token", res.data.sessionToken);
          setSuccessOpen(true);
          setAmount(null);
          setLoading(false);
        });
    } catch (e) {
      setLoading(false);
    }
  };

  return (
    <main
      style={{
        margin: "30px",
        textAlign: "center",
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
      }}
    >
      <Image src={EWalletPic} alt="e-wallets" loading="lazy" width={300} />
      <Typography variant="subtitle1">
        Valance: {user && user.valance}
      </Typography>
      <TextField label="Amount" onChange={handleAmount} />
      <LoadingButton
        color="primary"
        onClick={handleTopUp}
        loading={loading}
        loadingPosition="center"
        endIcon={<SendIcon />}
        variant="contained"
      >
        <span>TopUp!!</span>
      </LoadingButton>
      <CustomSnackBar
        open={errOpen}
        setOpen={setErrOpen}
        severity={Severity.ERROR}
        text={SnackBarTextErr}
      />
      <CustomSnackBar
        open={successOpen}
        setOpen={setSuccessOpen}
        severity={Severity.SUCCESS}
        text={SnackBarTextSuccess}
      />
    </main>
  );
};

export default TopUp;
