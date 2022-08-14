package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.LeagueAdapter;
import xavier.ricardo.paulobet.adapters.ScoresAdapter;
import xavier.ricardo.paulobet.model.GetLeaguesResponse;
import xavier.ricardo.paulobet.model.ScoresResponse;
import xavier.ricardo.paulobet.tasks.ScoresTask;

public class ScoresActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_scores);
        Intent intent = getIntent();
        String league = intent.getStringExtra("league");

        new ScoresTask(this, league).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            ScoresResponse scoresResponse = gson.fromJson(body, ScoresResponse.class);
            ListView lvScores = findViewById(R.id.lvScores);
            ScoresAdapter adapter = new ScoresAdapter(this, scoresResponse.getScores());
            lvScores.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }

    }
}