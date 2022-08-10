package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;
import java.util.List;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.LeagueAdapter;
import xavier.ricardo.paulobet.model.GetLeaguesResponse;
import xavier.ricardo.paulobet.model.League;
import xavier.ricardo.paulobet.tasks.GetLeaguesTask;

public class SelectLeagueActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_select_league);
        Intent intent = getIntent();
        String user = intent.getStringExtra("user");
        new GetLeaguesTask(this, user).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            GetLeaguesResponse getLeaguesResponse = gson.fromJson(body, GetLeaguesResponse.class);
            ListView lvLeagues = (ListView) findViewById(R.id.lvLeagues);
            LeagueAdapter adapter = new LeagueAdapter(this, getLeaguesResponse.getLeagues());
            lvLeagues.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }
}