package cmpln

import(
    "net/http"
    "strings"
    "strconv"
)



func HTTPDeletePost(w http.ResponseWriter, r *http.Request) {
    splitURL := strings.Split(r.URL.Path, "/")
    
    id, err := strconv.ParseInt( splitURL[len(splitURL) - 1], 10, 0) 
    
    if err != nil {
        http.Error(w, "Id not in URL", http.StatusInternalServerError)
        return        
    }

    if _, err := DeletePost(id); err != nil {
        http.Error(w, "Deletion not possible", http.StatusInternalServerError)
        return
    }

    // apparently used if DELETE method is used. dont really know why but seems to be best practise.
    w.WriteHeader(http.StatusNoContent)
}
