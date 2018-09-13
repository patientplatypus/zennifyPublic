<template>
  <div class="BlogFile">
    <br/>
    <h2>
      Blog
    </h2>
    <h3>
      September 12, 2018
    </h3>
    <p>
      Morning - continue to work on Cassandra mail - retrieval.
    </p>
    <p>
      Evening - can send headers and body in email and store, work on getting body on message click and more ui work tomorrow.
    </p>
    <br/>
    <h3>
      September 11, 2018
    </h3>
    <p>
      Morning - currently saving canvas passes data back and forth through IPFS correctly, saving the blob is tricky however. Moving on for now to work on Cassandra/profile explorer.
    </p>
    <p>
      Evening - modified dockerfile to get Cassandra to add appropriate keyspace, started on making frontend ui for messages, and have write message handler working (I think). Work on retrieving messages tomorrow.
    </p>
    <br/>
    <h3>
      September 10, 2018
    </h3>
    <p>
      Morning - work on getting canvas to perpetuate through IPFS, then work on messaging, then Cassandra.
    </p>
    <p>
      Evening - Got IPFS to perpetuate (set canvas tomorrow morning), Cassandra database instantiated. Work on pulling users into an "others" page, so I can then see profiles and make a messaging service with Cassandra.
    </p>
    <br/>
    <h3>
      September 9, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - no progress today.
    </p>
    <br/>
    <h3>
      September 8, 2018
    </h3>
    <p>
      Morning - get canvas color/size change operational then work on blockchain.
    </p>
    <p>
      Evening - got canvas to where I wanted, got WASM set up for rust. Tomorrow I need to look into following the blockchain tutorials and possible setting up ZEIT/NOW.
    </p>
    <br/>
    <h3>
      September 7, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - worked a bit on the canvas - more tomorrow.
    </p>
    <br/>
    <h3>
      September 6, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - boxes put on profile page with what should go there. Got IPFS storage of messages to work - done for now, need to later handle deletion of messages, a max number of messages (lucky 13 maybe?), better UI, also allow a user to send to another user (need to have a public page as well). Going on to other rough cuts of the profile page tomorrow. 
    </p>
    <br/>
    <h3>
      September 5, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - fixed minor bug in money section, file download now works (needs some tweaking - set user email as part of name). For the rough cut of the app probably wont add either auth or a profile postgres table although they would be useful. Continue working on profile tomorrow, specifically filling out some of the page UI. What other things go in a profile page?
    </p>
    <br/>
    <h3>
      September 4, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - File upload now works (needs authentication, tag to user account, frontendyness). Hot reload activated and confirmed works (realize). Work on profile page tomorrow.
    </p>
    <br/>
    <h3>
      September 3, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - basic functionality for the chat room is complete. Needs UI work and also need to check on why I am getting memory leaks (open github issue). Tomorrow work on health stuff or UI ~ something fun and not backend-y.
    </p>
    <br/>
    <h3>
      September 2, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - IPFS fell through this morning/last night. Chat is now functional, talk page scaffolded, channels partly made (subchannels less so). Work on that tomorrow.
    </p>
    <br/>
    <h3>
      September 1, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - two SO posts outstanding. It looks like unless someone has a solution IPFS will not work. Added recurring utility handler in backend,and continued to build out post routes for money page.
    </p>
    <br/>
    <h3>
      August 31, 2018
    </h3>
    <p>
      Morning - Began work, nothing to report.
    </p>
    <p>
      Evening - Other work on hold to get IPFS pubsub chat up and running in tasks. Put handling in Vuex and flesh out chatroom with hashtags tomorrow, then go back to work on profit/expense. BTW setimmediate transpiler probably no longer necessary as no longer using OrbitDB.
    </p>
    <br/>
    <h3>
      August 30, 2018
    </h3>
    <p>
      Morning - Began work, nothing to report.
    </p>
    <p>
      Evening - put in graph for expense recurrence, started on profit scaffolding, started on psql database hookups on profit. Finish that tomorrow, work on scaffolding, add in other submit hooks.
    </p>
    <br/> 
    <h3>
      August 29, 2018
    </h3>
    <p>
      Morning - begin working on recurrence equation and subsequent recurrence graphs. Want to do today as well: Let's get started on an actual README!
    </p>
    <p>
      Evening - got a service worker to appropriately munge dates - finish that up tomorrow and then put in the graph for the recurrence.
    </p>
    <br/>
    <h3>
      August 28, 2018
    </h3>
    <p>
      Morning - began work, nothing to report.
    </p>
    <p>
      Evening - dropdown stutter fixed, calendar recurrence implemented (however, will not export values; need to make a separate function for that for the recurring graphs), most of the user flow for expenses done except for graphs for recurrence/recurrence equation outside of graph. Work on those tomorrow.
    </p>
    <br/>
    <h3>
      August 27, 2018
    </h3>
    <p>
      Morning - working on making new loan graph.
    </p>
    <p>
      Evening - finished working on rough cut of new loan in expenses for the moment. Graphs are up and calendars are now properly formatted. Replicate work in the profits section tomorrow and then go and backfill the other options for both.
    </p>
    <br/>
    <h3>
      August 26, 2018
    </h3>
    <p>
      Morning - Moved calendar, added sliderbar css. Finishing up new loan user flow.
    </p>
    <p>
      Evening - got everything done except submit button and mixed chart. Do those tomorrow. Recurring calendar will also require new functionality for highlighting and returning multiple days.
    </p>
    <br/>
    <h3>
      August 25, 2018
    </h3>
    <p>
      Morning - started work; nothing to report.
    </p>
    <p>
      Evening - Finished drop downs and mostly done with calendar. Tomorrow I'm going to finish up the calendar with some interactivity for date picking, put that in its own frame, hopefully get the rest of the data shape organized, and start working on the preview graph.
    </p>
    <br/>
    <h3>
      August 24, 2018
    </h3>
    <p>
      Morning - starting work on getting the UI for the me page scaffolded.
    </p>
    <p>
      Evening - Have the beginnings of the UI scaffolding started with user/money. Have the selectBox 90% styled the way I want. Tomorrow I will finish that, finish the cost/profit user flow and data integration UI for money.
    </p>
    <br/>
    <h3>
      August 23, 2018
    </h3>
    <p>
      Morning - going to work on completing the home page today.
    </p>
    <p>
      Evening - turtle info button took much longer than anticipated. managed to get that done as well as integrate charts - quash minor bug on the backend for autologin. Tomorrow will scaffold rough UI for the "me" page. Goal should be to drive the database design around UI; see scope of the things i want to put in and then work towards that.
    </p>
    <br/>
    <h3>
      August 22, 2018
    </h3>
    <p>
      Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
    </p>
  </div>
</template>

<script>
export default {
  name: 'BlogFile',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .BlogFile{
    /* overflow: hidden;
    overflow-y: auto; */
  }
</style>
