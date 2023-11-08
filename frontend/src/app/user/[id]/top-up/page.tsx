"use client";
import React, { useState, useEffect } from "react";
import { Button, TextField, Typography } from "@mui/material";
import Image from "next/image";
import CustomSnackBar, { Severity } from "@/components/SnackBar";
import EWalletPic from "../../../../../public/e-wallet.png";
import { SnackBarTextErr, SnackBarTextSuccess } from "./text";
import { useAuth } from "@/app/userContext";
import axios from "@/api/axiosConfig";
import { AxiosResponse } from "axios";

const TopUp = () => {
  const { user, topUp } = useAuth();
  const [amount, setAmount] = useState("");
  const [errOpen, setErrOpen] = useState<boolean>(false);
  const [successOpen, setSuccessOpen] = useState<boolean>(false);

  const handleTopUp = () => {
    // Here you can add the logic to top up the user's account

    if (!amount) return;
    if (!Number(amount)) {
      setErrOpen(true);
    }
    if (!user) return;
    axios
      .post("/top-up", { valance: Number(amount), id: user.id })
      .then((res: AxiosResponse) => {
        if (res.status !== 200) return;
        topUp(user, Number(amount));
        localStorage.setItem("session_token", res.data.sessionToken);
        setSuccessOpen(true);
        setAmount("");
      });
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
      <Typography>This is Top up page</Typography>
      <Image src={EWalletPic} alt="e-wallets" loading="lazy" width={300} />
      <Typography>Valance: {user && user.valance}</Typography>
      <TextField
        label="Amount"
        value={amount}
        onChange={(e) => setAmount(e.target.value)}
        error={Number(amount) ? false : !amount ? false : true}
      />
      <Button onClick={handleTopUp}>Top Up</Button>
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
