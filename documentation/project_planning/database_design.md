``` mermaid
erDiagram
    
    Posts {
        uuid post_id PK "UUID"
        string creator FK "User who created this post"
        string title "Title of this post"
        string description
        string image_url "URL to the image posted"
    }
    
    Comments {
        int comment_id PK "Id of the comment"
        string text
        uuid post_id
        string username
    }
    
    Post_Reactions {
        uuid post_id
        int reaction_id 
    }
    
    Comment_Reactions {
        uuid commend_id
        int reaction_id 
    }
    
    Reactions {
        int reaction_id PK "Id of the reaction"
        string emoticon
        string username
    }
      
    Users {
        string username PK "The username"
        string profile_image "Profile image of user"
        string user_email "Email of the user"
    }
   
    Posts ||--|| Users : ""
    Comments ||--|| Users : ""
    Reactions ||--|| Users : ""
    Posts ||--|| Comments : ""
    Posts ||--|| Post_Reactions : ""
    Comments ||--|| Comment_Reactions : ""
    Post_Reactions ||--|| Reactions : ""
    Comment_Reactions ||--|| Reactions : ""
```