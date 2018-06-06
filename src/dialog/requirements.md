### My dialog project requirements

- ####  Login

    -  User can login through gmail, facebook or even through mobile number.

    -  User can post or view all dialogs as a guest user.

    -  Upon posting any Dialog as a guest user, it prompts for mobile number or email

- ####  View dialog(s)

    - Application displays top 10 (X number) of dialogs of the day.
    - User can search by language
    - Upon language search those top 10 (X number) dialogs to be posted
    - User can mark any dialog as his favorite dialog
    - User can type and search for any specific dialog by typing keywords in the search bar
    - Each dialog will have number of voice, video and selfies avaialable.
    - Upon selecting voice , video or selfie icons it shows top 10 voice, selfie and videos for that dialog.


- ####  Post dialog  

    - User can post his own dialog. Language and preffered actors can add for his dialog.
    - User can post his voice version or video version, selfie version for his or any dialog(s) in the system.
        - Voice version means ,telling any existing dialog in User's tone and post it
        - Selfie version means, take a selfie for any dialog and post it.
        - Video version means, take a video for existing or own dialog(s) and post it.

    - ##### API

        - /dialog : To post a dialog

        - /dialog/{dialog_id}/voice :   To post a voice for a dialog

        - /dialog/{dialog_id}/selfie : To post a selfie for a dialog

        - /dialog/{dialog_id}/video : To post a video for a dialog

- ####  Delete dialog
    
    - User can delete/archive his dialog from the system.

    - ##### API

        - /dialog/{dialog_id} : To delete a dialog. It is a hard delete.

- ####  Update dialog

    - User can update his dialog until no voice ,selfie or videos are posted by other users.

    - User can update language of and preffered artists for his dialog at any point of time.

     - ##### API

        - /dialog/{dialog_id} : To update a dialog.

- ### Like or Dislike dialog

    - User can like or dislike any dialog, voice , video or selfie.

    - ##### API

        - /dialog/{dialog_id}/like  : Put request to like it

        - /dialog/{dialog_id}/dislike : Put request to dislike it

- ####  Favorites

    - All user favorites are listed and displayed in favorites section. At first top 10 (X number) will be displayed can load another 10 (x number ) later.
    
    - User can add any dialog as his favorite

    - User can remove any dialog from his favorites at any point of time.

     - ##### API

        - /favorite/{user}/{count}/{offset}  : Get all favorites dialogs with given count and offset

        - /favorite/{user}/{dialog_id}/add : Put add a dialog to favorites

        - /favorite/{user}/{dialog_id}/remove : Remove a dialog from favorites.

- ### Actions

    - All user likes , dislikes are listed and displayed in user actions section. At first top 10 (X number) will be displayed can load another 10 (x number ) later.
    - User can clear this section, that permanently deletes from that section.

    - ##### API

        - /actions/{user}/{count}/{offset} : Get all user actions
       
        - /actions/{user}/clear : Clear all favorites for an user

- ### Notifications

    - User can opt for notifications for any language dialogs or any dialog voice , video or selfie reply.
