import React, {useState} from "react";
import {Theme} from "@material-ui/core";
import {makeStyles} from "@material-ui/core/styles";
import Dialog from "@material-ui/core/Dialog";
import PageContainer from "./PageContainer";
import PageToolbar from "./PageToolbar";
import CloseIcon from "@material-ui/icons/Close";
import IconButton from "@material-ui/core/IconButton";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Card from "@material-ui/core/Card";
import {banks} from "./banks";
import Slide from "@material-ui/core/Slide";
import Button from "@material-ui/core/Button";
import connectArrows from "../assets/connect-arrows.svg";
import Paper from "@material-ui/core/Paper";
import Accordion from "@material-ui/core/Accordion";
import AccordionSummary from "@material-ui/core/AccordionSummary";
import AccordionDetails from "@material-ui/core/AccordionDetails";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import financrooIcon from "../assets/banks/financroo-icon.svg";

const useStyles = makeStyles((theme: Theme) => ({
  cardRoot: {
    height: 116,
    padding: '0 16px',
    display: 'flex',
    '&:hover': {
      cursor: 'pointer'
    }
  },
  footer: {
    display: 'flex',
    position: 'fixed',
    justifyContent: 'center',
    alignItems: 'center',
    bottom: 0,
    height: 96,
    width: '100%',
    background: '#fff',
    borderTop: '1px solid #ECECEC'
  }
}));

export default ({onAllowAccess, onClose}) => {
  const classes = useStyles();
  const [selected, setSelected] = useState<any | null>(null);

  return (
    <Dialog open={true} fullScreen>
      <PageToolbar>
        {!selected && (
          <IconButton edge="start" color="inherit" aria-label="close" onClick={onClose}>
            <CloseIcon/>
          </IconButton>
        )}
        {selected && (
          <img src={selected.logo}/>
        )}
      </PageToolbar>
      <PageContainer style={{marginBottom: 112}} withBackground>
        {!selected && (
          <Grid container justify={'center'} style={{marginTop: 64}}>
            <Grid item xs={12} sm={8} md={6} style={{textAlign: 'center'}}>
              <Typography color={'primary'} variant={'h2'} style={{marginTop: 24, fontSize: 28}}>Connect your accounts</Typography>
              <Typography variant={'body1'} style={{marginTop: 16}}>By connecting your bank, bills and credit cards, you allow us to help you
                uncover
                insights that can improve your financial well being</Typography>
              <Grid container style={{marginTop: 48}} spacing={3}>
                {banks.map((bank) => (
                  <Grid item xs={6} sm={4}>
                    <Card className={classes.cardRoot} onClick={() => setSelected(bank)}>
                      <img src={bank.logo} style={{width: '100%'}}/>
                    </Card>
                  </Grid>
                ))}
              </Grid>
            </Grid>
          </Grid>
        )}
        <Slide direction="left" in={selected} mountOnEnter unmountOnExit exit={false}>
          <div>
            <Grid container justify={'center'} style={{marginTop: 64}}>
              <Grid item xs={12} sm={8} md={6} style={{textAlign: 'center'}}>
                <Typography color={'primary'} variant={'h2'} style={{marginTop: 24, fontSize: 28}}>Requested access</Typography>
                <Typography variant={'body1'} style={{marginTop: 16}}>In order to use this service, Financroo needs to access the following
                  information from your account service provider.</Typography>
                <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center', marginTop: 32}}>
                  <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#FCFCFF', width: 80, height: 80, borderRadius: '50%', border: '1px solid rgb(236 236 236)', marginRight: 16}}>
                    <img src={financrooIcon} style={{width: '60%'}}/>
                  </div>
                  <img src={connectArrows} style={{marginRight: 16}}/>
                  <div style={{display: 'flex', alignItems: 'center', justifyContent: 'center', background: '#FCFCFF', width: 80, height: 80, borderRadius: '50%', border: '1px solid rgb(236 236 236)'}}>
                    <img src={selected?.icon || selected?.logo} style={{width: '60%'}}/>
                  </div>
                </div>
                <Paper style={{marginTop: 32, padding: 16, textAlign: 'left'}}>
                  <Typography variant={'h4'} style={{fontSize: 16, marginBottom: 24}}>What we need you to share</Typography>
                  {[1, 2, 3, 4, 5].map(scope => (
                      <Accordion elevation={0}>
                        <AccordionSummary
                          expandIcon={<ExpandMoreIcon/>}
                          aria-controls="panel1a-content"
                          id="panel1a-header"
                        >
                          <Typography>Accordion {scope}</Typography>
                        </AccordionSummary>
                        <AccordionDetails>
                          <Typography>
                            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse malesuada lacus ex,
                            sit amet blandit leo lobortis eget.
                          </Typography>
                        </AccordionDetails>
                      </Accordion>
                    )
                  )}
                </Paper>
              </Grid>
            </Grid>
          </div>
        </Slide>
      </PageContainer>
      {selected && (
        <div className={classes.footer}>
          <div>
            <Button size={'large'} variant={'outlined'} onClick={() => setSelected(null)}>Cancel</Button>
            <Button size={'large'} variant={'contained'} color={'secondary'} style={{marginLeft: 16}} onClick={() => onAllowAccess({"account_id":"newElooAccountId","description":"newElooDescription","identification":"newElooIdentification","name":"newElooName"})}>Allow access</Button>
          </div>
        </div>
      )}
    </Dialog>
  )
};
