import React from "react";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Avatar,
    Button,
    Divider,
    Grid,
    Theme,
    Typography
} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import ExpandMoreIcon from "@material-ui/icons/ExpandMore";
import {permissionNameDescriptionMap} from "./permissionNameDescriptionMap";

const useStyles = makeStyles((theme: Theme) => ({
    expandIcon: {
        position: 'absolute', right: 32, top: 24, color: "#006580"
    }
}));

export default function ApplicationCard({consent, onRevokeClick}) {
    const classes = useStyles();

    return (
        <Accordion elevation={3}>
            <AccordionSummary
                classes={{expandIcon: classes.expandIcon}}
                expandIcon={<Avatar><ExpandMoreIcon/></Avatar>}
                aria-controls="panel1a-content"
                id="panel1a-header"
                style={{padding: "24px 32px"}}
            >
                <div>
                    <Avatar variant={"square"} style={{
                        background: "#D5E2E5",
                        color: "#006580"
                    }}>{consent?.name[0]?.toUpperCase()}</Avatar>
                </div>
                <div style={{width: "100%"}}>
                    <div style={{display: "flex", alignItems: "center"}}>
                        <Typography variant="h4"
                                    style={{marginLeft: 24}}>{consent.name}</Typography>
                    </div>
                    <Divider style={{margin: "24px 0 24px 24px"}}/>
                    <div style={{
                        display: "flex",
                        justifyContent: "space-between",
                        marginLeft: 24
                    }}>
                        <div>
                            <Typography
                                style={{color: "#626576", fontSize: 16}}><strong>Consent
                                validation</strong></Typography>
                            <div style={{display: "flex", marginTop: 16}}>
                                <Typography>Granted on</Typography>
                                <div style={{
                                    background: "#D5E2E5",
                                    color: "#006580",
                                    fontSize: 12,
                                    padding: 4,
                                    fontWeight: 500,
                                    marginLeft: 8
                                }}
                                >{consent.creation_date_time.split('T')[0]}
                                </div>
                                <Typography style={{marginLeft: 36}}>Valid
                                    until</Typography>
                                <div style={{
                                    background: "#D5E2E5",
                                    color: "#006580",
                                    fontSize: 12,
                                    padding: 4,
                                    fontWeight: 500,
                                    marginLeft: 8
                                }}
                                >{consent.expiration_date_time.split('T')[0]}
                                </div>
                            </div>
                        </div>
                        <Button
                            style={{color: "#DC1B37", textTransform: 'none'}}
                            onClick={e => {
                                e.stopPropagation();
                                onRevokeClick(consent);
                            }}
                        >Revoke Access
                        </Button>
                    </div>
                    <div style={{marginLeft: 24, marginTop: 32}}>
                        <Button
                            variant={"outlined"}
                            color={"primary"}
                            style={{textTransform: 'none'}}
                            onClick={e => {
                                e.stopPropagation();
                                window.open(consent.client_uri, "_blank")
                            }
                            }>View
                            information</Button>
                    </div>
                </div>
            </AccordionSummary>
            <AccordionDetails style={{
                flexDirection: "column",
                background: "#F4F7F8",
                padding: 0
            }}>
                {consent.account_ids.map(accountId => (
                    <div style={{paddingBottom: 32, borderLeft: "6px solid #006580",}}
                         key={accountId}>
                        <div style={{
                            background: "#D3E5EA",
                            height: 72,
                            color: "#006580",
                            paddingLeft: 88,
                            display: "flex",
                            alignItems: "center",
                            textTransform: 'uppercase'
                        }}>
                            <Typography>
                                <strong>Account ID: {accountId}</strong>
                            </Typography>
                        </div>
                        <div style={{paddingLeft: 88, color: "#006580"}}>
                            <Typography style={{
                                marginTop: 24,
                                borderBottom: "1px solid #ECECEC",
                                paddingBottom: 16,
                                marginBottom: 24,
                                textTransform: 'uppercase'
                            }}><strong>Granted permissions:</strong></Typography>
                            <Grid container spacing={3}>
                                {consent.permissions.map(permission => (
                                    <>
                                        <Grid item xs={4}
                                              style={{
                                                  textAlign: "left",
                                                  wordWrap: "break-word"
                                              }}>
                                            <Typography><strong>{permission}</strong></Typography>
                                        </Grid>
                                        <Grid item xs={8}
                                              style={{
                                                  textAlign: "left",
                                                  wordWrap: "break-word"
                                              }}>
                                            <Typography
                                                variant={"caption"}>{permissionNameDescriptionMap[permission]}</Typography>
                                        </Grid>
                                    </>
                                ))}
                            </Grid>
                        </div>
                    </div>
                ))}
            </AccordionDetails>
        </Accordion>
    )
};
