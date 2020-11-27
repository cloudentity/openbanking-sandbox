import React, {useEffect, useState} from "react";
import PageToolbar from "./PageToolbar";
import MenuIcon from "@material-ui/icons/Menu";
import IconButton from "@material-ui/core/IconButton";
import Progress from "./Progress";
import Tab from "@material-ui/core/Tab";
import Hidden from "@material-ui/core/Hidden";
import Tabs from "@material-ui/core/Tabs";
import {
    Accordion,
    AccordionDetails,
    AccordionSummary,
    Avatar,
    Button,
    Container,
    Grid,
    Theme,
    Typography
} from "@material-ui/core";
import {logout} from "./AuthPage";
import {makeStyles} from "@material-ui/core/styles";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import {api} from "../api/api";

const useStyles = makeStyles((theme: Theme) => ({
    indicator: {
        backgroundColor: '#fff',
    }
}));

export default function Dashboard({authorizationServerURL, authorizationServerId, tenantId}) {
    const [isProgress, setProgress] = useState(true);
    const [consents, setConsents] = useState<any>([]);
    const classes = useStyles();

    useEffect(() => {
        setProgress(true);
        api.getConsents()
            .then(res => setConsents(res))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }, []);

    const handleRemove = id => {
        setProgress(true);
        api.deleteConsent({id})
            .then(api.getConsents)
            .then(res => setConsents(res))
            .catch(err => console.log(err))
            .finally(() => setProgress(false));
    }

    return (
        <div>
            <PageToolbar>
                <Hidden mdUp>
                    <IconButton edge="start" color="inherit" aria-label="menu">
                        <MenuIcon/>
                    </IconButton>
                </Hidden>
                <Hidden smDown>
                    <Tabs
                        value={'connected-applications'}
                        onChange={() => {
                        }}
                        classes={{
                            indicator: classes.indicator
                        }}
                        aria-label="menu tabs"
                        style={{height: 64}}
                    >
                        <Tab label="Get started" value={'get-started'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="APIs" value={'apis'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Help Center" value={'help-center'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Connected Applications" value={'connected-applications'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                        <Tab label="Monitoring" value={'monitoring'}
                             style={{height: 64, textTransform: "none", minWidth: "unset"}}/>
                    </Tabs>
                </Hidden>
                <div style={{flex: 1}}/>
                <Button
                    style={{color: '#fff'}}
                    onClick={() => logout(authorizationServerURL, tenantId, authorizationServerId)}>Logout</Button>
            </PageToolbar>
            <div style={{marginTop: 64, position: "relative"}}>
                {isProgress && <Progress/>}

                {!isProgress && (
                    <>
                        <div style={{
                            background: "#F7F7F7",
                            height: 148,
                            display: 'flex',
                            alignItems: 'center'
                        }}>
                            <Container>
                                <Typography variant={"h3"} color={"primary"}>Connected Applications</Typography>
                            </Container>
                        </div>
                        <Container style={{marginTop: 64}}>
                            <Grid container justify={"center"}>
                                <Grid item xs={8}>
                                    {consents.length === 0 && (
                                        <Typography>No connected applications</Typography>
                                    )}
                                    {consents.map(consent => (
                                        <Accordion>
                                            <AccordionSummary
                                                expandIcon={<ExpandMoreIcon/>}
                                                aria-controls="panel1a-content"
                                                id="panel1a-header"
                                            >
                                                <div style={{
                                                    display: "flex",
                                                    justifyContent: "space-between",
                                                    alignItems: "center",
                                                    width: '100%'
                                                }}>
                                                    <div style={{display: "flex", alignItems: "center"}}>
                                                        <Avatar>{consent.name[0]}</Avatar>
                                                        <Typography style={{marginLeft: 16}}>{consent.name}</Typography>
                                                    </div>
                                                    <div style={{flex: 1}}/>
                                                    <Button variant={"contained"} color={"primary"} onClick={e => {
                                                        e.stopPropagation();
                                                        handleRemove(consent.id);
                                                    }}>Revoke</Button>
                                                </div>
                                            </AccordionSummary>
                                            <AccordionDetails style={{flexDirection: "column", background: "#fcfcfc"}}>
                                                {consent.accounts.map(account => (
                                                    <div style={{marginBottom: 32}}>
                                                        <Typography>
                                                            <strong>Account ID: {account.accountId}</strong>
                                                        </Typography>
                                                        <Typography style={{
                                                            marginTop: 16,
                                                            borderBottom: "1px solid #ccc",
                                                            paddingBottom: 8,
                                                            marginBottom: 16
                                                        }}>Granted permissions:</Typography>
                                                        <Grid container spacing={2}>
                                                            {account.permissions.map(permission => (
                                                                <>
                                                                    <Grid item xs={4}
                                                                          wrap={"wrap"}
                                                                          style={{
                                                                              textAlign: "right",
                                                                              wordWrap: "break-word"
                                                                          }}>
                                                                        <Typography><strong>{permission.name}</strong></Typography>
                                                                    </Grid>
                                                                    <Grid item xs={8}
                                                                          style={{
                                                                              textAlign: "left",
                                                                              wordWrap: "break-word"
                                                                          }}>
                                                                        <Typography>{permission.description}</Typography>
                                                                    </Grid>
                                                                </>
                                                            ))}

                                                        </Grid>
                                                    </div>
                                                ))}
                                            </AccordionDetails>
                                        </Accordion>
                                    ))}
                                </Grid>
                            </Grid>

                        </Container>
                    </>
                )}
            </div>
        </div>
    )
};
