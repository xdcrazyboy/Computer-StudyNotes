import java.util.ArrayList;
import java.util.Arrays;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello World!");
    }

    public List<List<Integer>> getFull(int[] arr){
        List<List<Integer>> res = new ArrayList<>();
        int len = arr.length;
        if(len == 0) {
            return res;
        }
        //默认值false
        boolean[] visited = new boolean[len];
        back(res, new ArrayList<>(), arr, visited);
        return res;
    }
    
    public List<List<Integer>> back(List<List<Integer>> res, List<Integer> ans, int[] source, boolean[] visited){
        if(ans.size() == source.length) {
            res.add(ans);
        }
        for (int i = 0; i < source.length; i++) {
            if(visited[i]){
               continue; 
            }
            ans.add(source[i]);
            visited[i] = true;
            back(res, ans, source, visited);
            visited[i] =  false;
            ans.remove(ans.size() - 1);
        }
    }
}
