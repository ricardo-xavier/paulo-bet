package xavier.ricardo.paulobet;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Intent;
import android.os.Bundle;
import android.widget.ListView;
import android.widget.Toast;

import com.google.gson.Gson;

import java.io.IOException;

import okhttp3.Response;
import xavier.ricardo.paulobet.adapters.BetsAdapter;
import xavier.ricardo.paulobet.adapters.ScoresAdapter;
import xavier.ricardo.paulobet.model.GetBetsResponse;
import xavier.ricardo.paulobet.model.GetScoresResponse;
import xavier.ricardo.paulobet.tasks.GetBetsTask;

public class GetBetsActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_get_bets);

        Intent intent = getIntent();
        String league = intent.getStringExtra("league");
        String user = intent.getStringExtra("user");

        new GetBetsTask(this, league, user).execute();
    }

    public void onTaskResponse(Response response) throws IOException {
        if (response.code() == 200) {
            String body = response.body().string();
            Gson gson = new Gson();
            GetBetsResponse betsResponse = gson.fromJson(body, GetBetsResponse.class);
            ListView lvBets = findViewById(R.id.lvBets);
            BetsAdapter adapter = new BetsAdapter(this, betsResponse.getBets());
            lvBets.setAdapter(adapter);
        } else {
            Toast.makeText(this, "ERRO status " + response.code(), Toast.LENGTH_LONG).show();
        }
    }
}