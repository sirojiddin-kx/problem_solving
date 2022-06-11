package main

func angryProfessor(k int32, a []int32) string {
    var (
        count int32
    )
    for _, arrival := range a {
        if arrival <= 0  {
            count += 1
        }
    }

    if count >= k {

        return "NO"
    } else {
        
        return "YES"
    }
    
}