import React from "react";
import ClientsList from "./ClientsList";
import {Typography} from "@material-ui/core";
import noAccountEmptyState from "./no-accounts-empty-state.svg";


export default ({clients, onRevokeClient, onRevokeConsent}) => {
    return (
        <>
            {clients.length === 0 && (
                <div style={{textAlign: "center", marginTop: 64}}>
                    <Typography variant={"h3"} style={{color: "#626576"}}>No connected
                        accounts</Typography>
                    <Typography style={{marginTop: 12, color: "#A0A3B5"}}>No connected
                        accounts to manage access</Typography>
                    <img src={noAccountEmptyState} style={{marginTop: 64}} alt={"empty state"}/>
                </div>
            )}
            {clients.length > 0 && <ClientsList clients={clients} onRevokeClient={onRevokeClient} onRevokeConsent={onRevokeConsent}/>}
        </>
    )
};
